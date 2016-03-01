// +build f10x_ld f10x_ld_vl f10x_md f10x_md_vl f10x_hd f10x_hd_vl f10x_xl

package rtc

import (
	"math"
	"rtos"
	"sync/atomic"
	"syscall"
	"time"

	"arch/cortexm/bitband"

	"stm32/hal/exti"
	"stm32/hal/irq"

	"stm32/hal/raw/bkp"
	"stm32/hal/raw/pwr"
	"stm32/hal/raw/rcc"
	"stm32/hal/raw/rtc"
)

// When 32768 Hz oscilator is used and preLog2 == 5 then:
// - rtos.Nanosec resolution will be 1/32768 s,
// - rtos.SleepUntil resoultion will be 1<<5/32768 s = 1/1024 s ≈ 1 ms,
// - the longest down time (RTC on battery) will be 1<<32/1024 s ≈ 48 days.
const (
	preLog2   = 5
	prescaler = 1 << preLog2

	maxSleepCnt  = 1 << 22
	maxSleepTick = maxSleepCnt << preLog2

	flagOK  = 0
	flagSet = 1
)

type globals struct {
	freqHz  uint
	cntExt  int32  // 16 bit RTC VCNT excension.
	lastISR uint32 // Last ISR  time using uint32(VCNT).
	status  bitband.Bits16
}

var g globals

func init() {
	status := rtcBackup{bkp.BKP}.Status()
	g.status = bitband.Alias16(status)
}

func setup(freqHz uint) {
	g.freqHz = freqHz

	RTC := rtc.RTC
	RCC := rcc.RCC
	PWR := pwr.PWR
	bkp := rtcBackup{bkp.BKP}

	const (
		mask = rcc.LSEON | rcc.RTCSEL | rcc.RTCEN
		cfg  = rcc.LSEON | rcc.RTCSEL_LSE | rcc.RTCEN
	)

	// Enable write access to backup domain.
	RCC.APB1ENR.SetBits(rcc.PWREN | rcc.BKPEN)
	_ = RCC.APB1ENR.Load()
	PWR.DBP().Set()
	RCC.APB1ENR.ClearBits(rcc.PWREN)

	if RCC.BDCR.Bits(mask) != cfg || g.status.Bit(flagOK).Load() == 0 {
		// RTC not initialized or in dirty state.

		// Reset backup domain and configure RTC clock source.
		RCC.BDRST().Set()
		RCC.BDRST().Clear()
		RCC.LSEON().Set()
		for RCC.LSERDY().Load() == 0 {
		}
		RCC.BDCR.StoreBits(mask, cfg)

		// Configure RTC prescaler.
		waitForSync(RTC)
		waitForWrite(RTC)
		RTC.CNF().Set() // Begin PRL configuration
		RTC.PRLL.Store(prescaler - 1)
		RTC.CNF().Clear() // Copy from APB to BKP domain.

		g.status.Bit(flagOK).Set()

		// Wait for complete before setup RTCALR interrupt.
		waitForWrite(RTC)
	} else {
		g.cntExt = int32(int16(bkp.CntExt().Load()))
		g.lastISR = bkp.LastISR().Load()
		if g.status.Bit(flagSet).Load() != 0 {
			sec := bkp.StartSec().Load()
			ns := int32(bkp.StartNanosec().Load())
			start := time.Unix(int64(sec), int64(ns))
			time.SetStart(start)
		}
	}
	// Wait for sync. Need in both cases: after reset (synchronise APB domain)
	// or after configuration (avoid reading bad DIVL).
	waitForSync(RTC)

	exti.RTCALR.EnableRiseTrig()
	exti.RTCALR.EnableInt()
	spnum := rtos.IRQPrioStep * rtos.IRQPrioNum
	rtos.IRQ(irq.RTCAlarm).SetPrio(rtos.IRQPrioLowest + spnum*3/4)
	rtos.IRQ(irq.RTCAlarm).Enable()

	// Force RTCISR to early handle possible overflow.
	exti.RTCALR.Trigger()

	syscall.SetSysTimer(nanosec, setWakeup)
}

// loadVCNT returns value of virtual counter that counts number of ticks of
// RTC input clock. Value of this virtual counter is calculated according to
// the formula:
//
//  VCNT = ((CNTH<<16 + CNTL)<<preLog2 + frac) & (prescaler<<32 - 1)
//
// where frac is calculated as follow:
//
//  frac = prescaler - (DIVL+1)&(prescaler-1)
//
// Only DIVL is used, so prescaler can not be greater than 0x10000.
//
// Thanks to this transformation, RTC interrupts are generated at right time.
// See example for Second, Overflow and Alarm(0-1) interrupts in table below:
//
//  CNT      DIV| VCNT>>5 VCNT&0x1f
//  ------------+--------------------
//  ffffffff 04 | ffffffff 1b
//  ffffffff 03 | ffffffff 1c
//  ffffffff 02 | ffffffff 1d
//  ffffffff 01 | ffffffff 1e
//  ffffffff 00 | ffffffff 1f
//  ffffffff 1f | 00000000 00 <- Second, Overflow, Alarm(0xffffffff)
//  00000000 1e | 00000000 01
//  00000000 1d | 00000000 02
//  00000000 1c | 00000000 03
//  00000000 1b | 00000000 04
//  00000000 1a | 00000000 05
//
func loadVCNT() int64 {
	RTC := rtc.RTC
	var (
		ch rtc.CNTH_Bits
		cl rtc.CNTL_Bits
		dl rtc.DIVL_Bits
	)
	ch = RTC.CNTH.Load()
	for {
		cl = RTC.CNTL.Load()
		for {
			dl = RTC.DIVL.Load()
			cl1 := RTC.CNTL.Load()
			if cl1 == cl {
				break
			}
			cl = cl1
		}
		ch1 := RTC.CNTH.Load()
		if ch1 == ch {
			break
		}
		ch = ch1
	}
	cnt := uint32(ch)<<16 | uint32(cl)
	frac := prescaler - (uint32(dl)+1)&(prescaler-1)
	return (int64(cnt)<<preLog2 + int64(frac)) & (prescaler<<32 - 1)
}

func isr() {
	exti.RTCALR.ClearPending()

	vcnt32 := uint32(loadVCNT() >> preLog2)
	if vcnt32 != g.lastISR {
		if vcnt32 < g.lastISR {
			cntext := g.cntExt + 1
			bkp := rtcBackup{bkp.BKP}
			g.status.Bit(flagOK).Clear()
			bkp.CntExt().Store(uint16(cntext))
			bkp.LastISR().Store(vcnt32)
			g.status.Bit(flagOK).Set()
			atomic.StoreInt32(&g.cntExt, cntext)
		} else {
			rtcBackup{bkp.BKP}.LastISR().Store(vcnt32)
		}
		atomic.StoreUint32(&g.lastISR, vcnt32)
	}
	syscall.SchedNext()
}

func loadTicks() int64 {
	irq.RTCAlarm.Disable()
	lastisr := atomic.LoadUint32(&g.lastISR)
	cntext := atomic.LoadInt32(&g.cntExt)
	vcnt := loadVCNT()
	irq.RTCAlarm.Enable()

	if uint32(vcnt>>preLog2) < lastisr {
		cntext++
	}
	return int64(cntext)<<(32+preLog2) | vcnt
}

// nanosec: see syscall.SetSysClock.
func nanosec() int64 {
	return ticktons(loadTicks())
}

// setWakeup: see syscall.SetSysClock.
func setWakeup(ns int64) {
	wkup := nstotick(ns)
	now := loadTicks()
	sleep := wkup - now
	switch {
	case sleep > maxSleepTick:
		wkup = now + maxSleepTick
	case sleep < 0:
		wkup = now
	}
	wkup = (wkup + prescaler/2) >> preLog2
	alr := uint32(wkup) - 1

	RTC := rtc.RTC
	waitForWrite(RTC)
	RTC.CNF().Set()
	RTC.ALRH.Store(rtc.ALRH_Bits(alr >> 16))
	RTC.ALRL.Store(rtc.ALRL_Bits(alr))
	RTC.CNF().Clear()

	if loadTicks()>>preLog2 >= wkup {
		// There is a chance that the alarm interrupt was not triggered.
		exti.RTCALR.Trigger()
	}
	/*print64("*sw* now:", now)
	print64(" wkup:", wkup)
	print32(" cnt1:", uint32(now>>rtcPreLog2))
	println32(" alr:", alr)*/
}

func setStartTime(t time.Time) {
	bkp := rtcBackup{bkp.BKP}
	if g.status.Bit(flagOK).Load() == 0 {
		return
	}
	sec := t.Unix()
	ns := t.Nanosecond()
	time.SetStart(t)
	g.status.Bit(flagOK).Clear()
	bkp.StartSec().Store(uint64(sec))
	bkp.StartNanosec().Store(uint32(ns))
	bkp.Status().Store(1<<flagSet | 1<<flagOK)
}

func status() (ok, set bool) {
	s := rtcBackup{bkp.BKP}.Status().Load()
	ok = s&(1<<flagOK) != 0
	set = s&(1<<flagSet) != 0
	return
}

func ticktons(tick int64) int64 {
	return int64(math.Muldiv(uint64(tick), 1e9, uint64(g.freqHz)))
}

func nstotick(ns int64) int64 {
	return int64(math.Muldiv(uint64(ns), uint64(g.freqHz), 1e9))
}