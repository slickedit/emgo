package tim

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type TIM_Periph struct {
	CR1   CR1
	_     uint16
	CR2   CR2
	_     uint16
	SMCR  SMCR
	_     uint16
	DIER  DIER
	_     uint16
	SR    SR
	_     uint16
	EGR   EGR
	_     uint16
	CCMR1 CCMR1
	_     uint16
	CCMR2 CCMR2
	_     uint16
	CCER  CCER
	_     uint16
	CNT   CNT
	_     uint16
	PSC   PSC
	_     uint16
	ARR   ARR
	_     uint16
	RCR   RCR
	_     uint16
	CCR1  CCR1
	_     uint16
	CCR2  CCR2
	_     uint16
	CCR3  CCR3
	_     uint16
	CCR4  CCR4
	_     uint16
	BDTR  BDTR
	_     uint16
	DCR   DCR
	_     uint16
	DMAR  DMAR
}

func (p *TIM_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var TIM2 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM2_BASE)))

//emgo:const
var TIM3 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM3_BASE)))

//emgo:const
var TIM4 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM4_BASE)))

//emgo:const
var TIM5 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM5_BASE)))

//emgo:const
var TIM6 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM6_BASE)))

//emgo:const
var TIM7 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM7_BASE)))

//emgo:const
var TIM12 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM12_BASE)))

//emgo:const
var TIM13 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM13_BASE)))

//emgo:const
var TIM14 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM14_BASE)))

//emgo:const
var TIM1 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM1_BASE)))

//emgo:const
var TIM8 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM8_BASE)))

//emgo:const
var TIM15 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM15_BASE)))

//emgo:const
var TIM16 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM16_BASE)))

//emgo:const
var TIM17 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM17_BASE)))

//emgo:const
var TIM9 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM9_BASE)))

//emgo:const
var TIM10 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM10_BASE)))

//emgo:const
var TIM11 = (*TIM_Periph)(unsafe.Pointer(uintptr(mmap.TIM11_BASE)))

type CR1_Bits uint16

func (b CR1_Bits) Field(mask CR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1_Bits) J(v int) CR1_Bits {
	return CR1_Bits(bits.Make32(v, uint32(mask)))
}

type CR1 struct{ mmio.U16 }

func (r *CR1) Bits(mask CR1_Bits) CR1_Bits { return CR1_Bits(r.U16.Bits(uint16(mask))) }
func (r *CR1) StoreBits(mask, b CR1_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CR1) SetBits(mask CR1_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CR1) ClearBits(mask CR1_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CR1) Load() CR1_Bits              { return CR1_Bits(r.U16.Load()) }
func (r *CR1) Store(b CR1_Bits)            { r.U16.Store(uint16(b)) }

type CR1_Mask struct{ mmio.UM16 }

func (rm CR1_Mask) Load() CR1_Bits   { return CR1_Bits(rm.UM16.Load()) }
func (rm CR1_Mask) Store(b CR1_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) CEN() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(CEN)}}
}

func (p *TIM_Periph) UDIS() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(UDIS)}}
}

func (p *TIM_Periph) URS() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(URS)}}
}

func (p *TIM_Periph) OPM() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(OPM)}}
}

func (p *TIM_Periph) DIR() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(DIR)}}
}

func (p *TIM_Periph) CMS() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(CMS)}}
}

func (p *TIM_Periph) ARPE() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(ARPE)}}
}

func (p *TIM_Periph) CKD() CR1_Mask {
	return CR1_Mask{mmio.UM16{&p.CR1.U16, uint16(CKD)}}
}

type CR2_Bits uint16

func (b CR2_Bits) Field(mask CR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2_Bits) J(v int) CR2_Bits {
	return CR2_Bits(bits.Make32(v, uint32(mask)))
}

type CR2 struct{ mmio.U16 }

func (r *CR2) Bits(mask CR2_Bits) CR2_Bits { return CR2_Bits(r.U16.Bits(uint16(mask))) }
func (r *CR2) StoreBits(mask, b CR2_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CR2) SetBits(mask CR2_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CR2) ClearBits(mask CR2_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CR2) Load() CR2_Bits              { return CR2_Bits(r.U16.Load()) }
func (r *CR2) Store(b CR2_Bits)            { r.U16.Store(uint16(b)) }

type CR2_Mask struct{ mmio.UM16 }

func (rm CR2_Mask) Load() CR2_Bits   { return CR2_Bits(rm.UM16.Load()) }
func (rm CR2_Mask) Store(b CR2_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) CCPC() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(CCPC)}}
}

func (p *TIM_Periph) CCUS() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(CCUS)}}
}

func (p *TIM_Periph) CCDS() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(CCDS)}}
}

func (p *TIM_Periph) MMS() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(MMS)}}
}

func (p *TIM_Periph) TI1S() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(TI1S)}}
}

func (p *TIM_Periph) OIS1() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(OIS1)}}
}

func (p *TIM_Periph) OIS1N() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(OIS1N)}}
}

func (p *TIM_Periph) OIS2() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(OIS2)}}
}

func (p *TIM_Periph) OIS2N() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(OIS2N)}}
}

func (p *TIM_Periph) OIS3() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(OIS3)}}
}

func (p *TIM_Periph) OIS3N() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(OIS3N)}}
}

func (p *TIM_Periph) OIS4() CR2_Mask {
	return CR2_Mask{mmio.UM16{&p.CR2.U16, uint16(OIS4)}}
}

type SMCR_Bits uint16

func (b SMCR_Bits) Field(mask SMCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMCR_Bits) J(v int) SMCR_Bits {
	return SMCR_Bits(bits.Make32(v, uint32(mask)))
}

type SMCR struct{ mmio.U16 }

func (r *SMCR) Bits(mask SMCR_Bits) SMCR_Bits { return SMCR_Bits(r.U16.Bits(uint16(mask))) }
func (r *SMCR) StoreBits(mask, b SMCR_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *SMCR) SetBits(mask SMCR_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *SMCR) ClearBits(mask SMCR_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *SMCR) Load() SMCR_Bits               { return SMCR_Bits(r.U16.Load()) }
func (r *SMCR) Store(b SMCR_Bits)             { r.U16.Store(uint16(b)) }

type SMCR_Mask struct{ mmio.UM16 }

func (rm SMCR_Mask) Load() SMCR_Bits   { return SMCR_Bits(rm.UM16.Load()) }
func (rm SMCR_Mask) Store(b SMCR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) SMS() SMCR_Mask {
	return SMCR_Mask{mmio.UM16{&p.SMCR.U16, uint16(SMS)}}
}

func (p *TIM_Periph) TS() SMCR_Mask {
	return SMCR_Mask{mmio.UM16{&p.SMCR.U16, uint16(TS)}}
}

func (p *TIM_Periph) MSM() SMCR_Mask {
	return SMCR_Mask{mmio.UM16{&p.SMCR.U16, uint16(MSM)}}
}

func (p *TIM_Periph) ETF() SMCR_Mask {
	return SMCR_Mask{mmio.UM16{&p.SMCR.U16, uint16(ETF)}}
}

func (p *TIM_Periph) ETPS() SMCR_Mask {
	return SMCR_Mask{mmio.UM16{&p.SMCR.U16, uint16(ETPS)}}
}

func (p *TIM_Periph) ECE() SMCR_Mask {
	return SMCR_Mask{mmio.UM16{&p.SMCR.U16, uint16(ECE)}}
}

func (p *TIM_Periph) ETP() SMCR_Mask {
	return SMCR_Mask{mmio.UM16{&p.SMCR.U16, uint16(ETP)}}
}

type DIER_Bits uint16

func (b DIER_Bits) Field(mask DIER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIER_Bits) J(v int) DIER_Bits {
	return DIER_Bits(bits.Make32(v, uint32(mask)))
}

type DIER struct{ mmio.U16 }

func (r *DIER) Bits(mask DIER_Bits) DIER_Bits { return DIER_Bits(r.U16.Bits(uint16(mask))) }
func (r *DIER) StoreBits(mask, b DIER_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DIER) SetBits(mask DIER_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DIER) ClearBits(mask DIER_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DIER) Load() DIER_Bits               { return DIER_Bits(r.U16.Load()) }
func (r *DIER) Store(b DIER_Bits)             { r.U16.Store(uint16(b)) }

type DIER_Mask struct{ mmio.UM16 }

func (rm DIER_Mask) Load() DIER_Bits   { return DIER_Bits(rm.UM16.Load()) }
func (rm DIER_Mask) Store(b DIER_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) UIE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(UIE)}}
}

func (p *TIM_Periph) CC1IE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(CC1IE)}}
}

func (p *TIM_Periph) CC2IE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(CC2IE)}}
}

func (p *TIM_Periph) CC3IE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(CC3IE)}}
}

func (p *TIM_Periph) CC4IE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(CC4IE)}}
}

func (p *TIM_Periph) COMIE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(COMIE)}}
}

func (p *TIM_Periph) TIE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(TIE)}}
}

func (p *TIM_Periph) BIE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(BIE)}}
}

func (p *TIM_Periph) UDE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(UDE)}}
}

func (p *TIM_Periph) CC1DE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(CC1DE)}}
}

func (p *TIM_Periph) CC2DE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(CC2DE)}}
}

func (p *TIM_Periph) CC3DE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(CC3DE)}}
}

func (p *TIM_Periph) CC4DE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(CC4DE)}}
}

func (p *TIM_Periph) COMDE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(COMDE)}}
}

func (p *TIM_Periph) TDE() DIER_Mask {
	return DIER_Mask{mmio.UM16{&p.DIER.U16, uint16(TDE)}}
}

type SR_Bits uint16

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U16 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U16.Bits(uint16(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U16.SetBits(uint16(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U16.ClearBits(uint16(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U16.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U16.Store(uint16(b)) }

type SR_Mask struct{ mmio.UM16 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM16.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) UIF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(UIF)}}
}

func (p *TIM_Periph) CC1IF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(CC1IF)}}
}

func (p *TIM_Periph) CC2IF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(CC2IF)}}
}

func (p *TIM_Periph) CC3IF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(CC3IF)}}
}

func (p *TIM_Periph) CC4IF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(CC4IF)}}
}

func (p *TIM_Periph) COMIF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(COMIF)}}
}

func (p *TIM_Periph) TIF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(TIF)}}
}

func (p *TIM_Periph) BIF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(BIF)}}
}

func (p *TIM_Periph) CC1OF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(CC1OF)}}
}

func (p *TIM_Periph) CC2OF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(CC2OF)}}
}

func (p *TIM_Periph) CC3OF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(CC3OF)}}
}

func (p *TIM_Periph) CC4OF() SR_Mask {
	return SR_Mask{mmio.UM16{&p.SR.U16, uint16(CC4OF)}}
}

type EGR_Bits uint16

func (b EGR_Bits) Field(mask EGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EGR_Bits) J(v int) EGR_Bits {
	return EGR_Bits(bits.Make32(v, uint32(mask)))
}

type EGR struct{ mmio.U16 }

func (r *EGR) Bits(mask EGR_Bits) EGR_Bits { return EGR_Bits(r.U16.Bits(uint16(mask))) }
func (r *EGR) StoreBits(mask, b EGR_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *EGR) SetBits(mask EGR_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *EGR) ClearBits(mask EGR_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *EGR) Load() EGR_Bits              { return EGR_Bits(r.U16.Load()) }
func (r *EGR) Store(b EGR_Bits)            { r.U16.Store(uint16(b)) }

type EGR_Mask struct{ mmio.UM16 }

func (rm EGR_Mask) Load() EGR_Bits   { return EGR_Bits(rm.UM16.Load()) }
func (rm EGR_Mask) Store(b EGR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) UG() EGR_Mask {
	return EGR_Mask{mmio.UM16{&p.EGR.U16, uint16(UG)}}
}

func (p *TIM_Periph) CC1G() EGR_Mask {
	return EGR_Mask{mmio.UM16{&p.EGR.U16, uint16(CC1G)}}
}

func (p *TIM_Periph) CC2G() EGR_Mask {
	return EGR_Mask{mmio.UM16{&p.EGR.U16, uint16(CC2G)}}
}

func (p *TIM_Periph) CC3G() EGR_Mask {
	return EGR_Mask{mmio.UM16{&p.EGR.U16, uint16(CC3G)}}
}

func (p *TIM_Periph) CC4G() EGR_Mask {
	return EGR_Mask{mmio.UM16{&p.EGR.U16, uint16(CC4G)}}
}

func (p *TIM_Periph) COMG() EGR_Mask {
	return EGR_Mask{mmio.UM16{&p.EGR.U16, uint16(COMG)}}
}

func (p *TIM_Periph) TG() EGR_Mask {
	return EGR_Mask{mmio.UM16{&p.EGR.U16, uint16(TG)}}
}

func (p *TIM_Periph) BG() EGR_Mask {
	return EGR_Mask{mmio.UM16{&p.EGR.U16, uint16(BG)}}
}

type CCMR1_Bits uint16

func (b CCMR1_Bits) Field(mask CCMR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCMR1_Bits) J(v int) CCMR1_Bits {
	return CCMR1_Bits(bits.Make32(v, uint32(mask)))
}

type CCMR1 struct{ mmio.U16 }

func (r *CCMR1) Bits(mask CCMR1_Bits) CCMR1_Bits { return CCMR1_Bits(r.U16.Bits(uint16(mask))) }
func (r *CCMR1) StoreBits(mask, b CCMR1_Bits)    { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CCMR1) SetBits(mask CCMR1_Bits)         { r.U16.SetBits(uint16(mask)) }
func (r *CCMR1) ClearBits(mask CCMR1_Bits)       { r.U16.ClearBits(uint16(mask)) }
func (r *CCMR1) Load() CCMR1_Bits                { return CCMR1_Bits(r.U16.Load()) }
func (r *CCMR1) Store(b CCMR1_Bits)              { r.U16.Store(uint16(b)) }

type CCMR1_Mask struct{ mmio.UM16 }

func (rm CCMR1_Mask) Load() CCMR1_Bits   { return CCMR1_Bits(rm.UM16.Load()) }
func (rm CCMR1_Mask) Store(b CCMR1_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) CC1S() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(CC1S)}}
}

func (p *TIM_Periph) OC1FE() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(OC1FE)}}
}

func (p *TIM_Periph) OC1PE() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(OC1PE)}}
}

func (p *TIM_Periph) OC1M() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(OC1M)}}
}

func (p *TIM_Periph) OC1CE() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(OC1CE)}}
}

func (p *TIM_Periph) CC2S() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(CC2S)}}
}

func (p *TIM_Periph) OC2FE() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(OC2FE)}}
}

func (p *TIM_Periph) OC2PE() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(OC2PE)}}
}

func (p *TIM_Periph) OC2M() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(OC2M)}}
}

func (p *TIM_Periph) OC2CE() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(OC2CE)}}
}

func (p *TIM_Periph) IC1PSC() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(IC1PSC)}}
}

func (p *TIM_Periph) IC1F() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(IC1F)}}
}

func (p *TIM_Periph) IC2PSC() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(IC2PSC)}}
}

func (p *TIM_Periph) IC2F() CCMR1_Mask {
	return CCMR1_Mask{mmio.UM16{&p.CCMR1.U16, uint16(IC2F)}}
}

type CCMR2_Bits uint16

func (b CCMR2_Bits) Field(mask CCMR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCMR2_Bits) J(v int) CCMR2_Bits {
	return CCMR2_Bits(bits.Make32(v, uint32(mask)))
}

type CCMR2 struct{ mmio.U16 }

func (r *CCMR2) Bits(mask CCMR2_Bits) CCMR2_Bits { return CCMR2_Bits(r.U16.Bits(uint16(mask))) }
func (r *CCMR2) StoreBits(mask, b CCMR2_Bits)    { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CCMR2) SetBits(mask CCMR2_Bits)         { r.U16.SetBits(uint16(mask)) }
func (r *CCMR2) ClearBits(mask CCMR2_Bits)       { r.U16.ClearBits(uint16(mask)) }
func (r *CCMR2) Load() CCMR2_Bits                { return CCMR2_Bits(r.U16.Load()) }
func (r *CCMR2) Store(b CCMR2_Bits)              { r.U16.Store(uint16(b)) }

type CCMR2_Mask struct{ mmio.UM16 }

func (rm CCMR2_Mask) Load() CCMR2_Bits   { return CCMR2_Bits(rm.UM16.Load()) }
func (rm CCMR2_Mask) Store(b CCMR2_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) CC3S() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(CC3S)}}
}

func (p *TIM_Periph) OC3FE() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(OC3FE)}}
}

func (p *TIM_Periph) OC3PE() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(OC3PE)}}
}

func (p *TIM_Periph) OC3M() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(OC3M)}}
}

func (p *TIM_Periph) OC3CE() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(OC3CE)}}
}

func (p *TIM_Periph) CC4S() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(CC4S)}}
}

func (p *TIM_Periph) OC4FE() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(OC4FE)}}
}

func (p *TIM_Periph) OC4PE() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(OC4PE)}}
}

func (p *TIM_Periph) OC4M() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(OC4M)}}
}

func (p *TIM_Periph) OC4CE() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(OC4CE)}}
}

func (p *TIM_Periph) IC3PSC() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(IC3PSC)}}
}

func (p *TIM_Periph) IC3F() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(IC3F)}}
}

func (p *TIM_Periph) IC4PSC() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(IC4PSC)}}
}

func (p *TIM_Periph) IC4F() CCMR2_Mask {
	return CCMR2_Mask{mmio.UM16{&p.CCMR2.U16, uint16(IC4F)}}
}

type CCER_Bits uint16

func (b CCER_Bits) Field(mask CCER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCER_Bits) J(v int) CCER_Bits {
	return CCER_Bits(bits.Make32(v, uint32(mask)))
}

type CCER struct{ mmio.U16 }

func (r *CCER) Bits(mask CCER_Bits) CCER_Bits { return CCER_Bits(r.U16.Bits(uint16(mask))) }
func (r *CCER) StoreBits(mask, b CCER_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CCER) SetBits(mask CCER_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *CCER) ClearBits(mask CCER_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *CCER) Load() CCER_Bits               { return CCER_Bits(r.U16.Load()) }
func (r *CCER) Store(b CCER_Bits)             { r.U16.Store(uint16(b)) }

type CCER_Mask struct{ mmio.UM16 }

func (rm CCER_Mask) Load() CCER_Bits   { return CCER_Bits(rm.UM16.Load()) }
func (rm CCER_Mask) Store(b CCER_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) CC1E() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC1E)}}
}

func (p *TIM_Periph) CC1P() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC1P)}}
}

func (p *TIM_Periph) CC1NE() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC1NE)}}
}

func (p *TIM_Periph) CC1NP() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC1NP)}}
}

func (p *TIM_Periph) CC2E() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC2E)}}
}

func (p *TIM_Periph) CC2P() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC2P)}}
}

func (p *TIM_Periph) CC2NE() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC2NE)}}
}

func (p *TIM_Periph) CC2NP() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC2NP)}}
}

func (p *TIM_Periph) CC3E() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC3E)}}
}

func (p *TIM_Periph) CC3P() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC3P)}}
}

func (p *TIM_Periph) CC3NE() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC3NE)}}
}

func (p *TIM_Periph) CC3NP() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC3NP)}}
}

func (p *TIM_Periph) CC4E() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC4E)}}
}

func (p *TIM_Periph) CC4P() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC4P)}}
}

func (p *TIM_Periph) CC4NP() CCER_Mask {
	return CCER_Mask{mmio.UM16{&p.CCER.U16, uint16(CC4NP)}}
}

type CNT_Bits uint16

func (b CNT_Bits) Field(mask CNT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CNT_Bits) J(v int) CNT_Bits {
	return CNT_Bits(bits.Make32(v, uint32(mask)))
}

type CNT struct{ mmio.U16 }

func (r *CNT) Bits(mask CNT_Bits) CNT_Bits { return CNT_Bits(r.U16.Bits(uint16(mask))) }
func (r *CNT) StoreBits(mask, b CNT_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CNT) SetBits(mask CNT_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CNT) ClearBits(mask CNT_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CNT) Load() CNT_Bits              { return CNT_Bits(r.U16.Load()) }
func (r *CNT) Store(b CNT_Bits)            { r.U16.Store(uint16(b)) }

type CNT_Mask struct{ mmio.UM16 }

func (rm CNT_Mask) Load() CNT_Bits   { return CNT_Bits(rm.UM16.Load()) }
func (rm CNT_Mask) Store(b CNT_Bits) { rm.UM16.Store(uint16(b)) }

type PSC_Bits uint16

func (b PSC_Bits) Field(mask PSC_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PSC_Bits) J(v int) PSC_Bits {
	return PSC_Bits(bits.Make32(v, uint32(mask)))
}

type PSC struct{ mmio.U16 }

func (r *PSC) Bits(mask PSC_Bits) PSC_Bits { return PSC_Bits(r.U16.Bits(uint16(mask))) }
func (r *PSC) StoreBits(mask, b PSC_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *PSC) SetBits(mask PSC_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *PSC) ClearBits(mask PSC_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *PSC) Load() PSC_Bits              { return PSC_Bits(r.U16.Load()) }
func (r *PSC) Store(b PSC_Bits)            { r.U16.Store(uint16(b)) }

type PSC_Mask struct{ mmio.UM16 }

func (rm PSC_Mask) Load() PSC_Bits   { return PSC_Bits(rm.UM16.Load()) }
func (rm PSC_Mask) Store(b PSC_Bits) { rm.UM16.Store(uint16(b)) }

type ARR_Bits uint16

func (b ARR_Bits) Field(mask ARR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ARR_Bits) J(v int) ARR_Bits {
	return ARR_Bits(bits.Make32(v, uint32(mask)))
}

type ARR struct{ mmio.U16 }

func (r *ARR) Bits(mask ARR_Bits) ARR_Bits { return ARR_Bits(r.U16.Bits(uint16(mask))) }
func (r *ARR) StoreBits(mask, b ARR_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *ARR) SetBits(mask ARR_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *ARR) ClearBits(mask ARR_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *ARR) Load() ARR_Bits              { return ARR_Bits(r.U16.Load()) }
func (r *ARR) Store(b ARR_Bits)            { r.U16.Store(uint16(b)) }

type ARR_Mask struct{ mmio.UM16 }

func (rm ARR_Mask) Load() ARR_Bits   { return ARR_Bits(rm.UM16.Load()) }
func (rm ARR_Mask) Store(b ARR_Bits) { rm.UM16.Store(uint16(b)) }

type RCR_Bits uint16

func (b RCR_Bits) Field(mask RCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RCR_Bits) J(v int) RCR_Bits {
	return RCR_Bits(bits.Make32(v, uint32(mask)))
}

type RCR struct{ mmio.U16 }

func (r *RCR) Bits(mask RCR_Bits) RCR_Bits { return RCR_Bits(r.U16.Bits(uint16(mask))) }
func (r *RCR) StoreBits(mask, b RCR_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RCR) SetBits(mask RCR_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *RCR) ClearBits(mask RCR_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *RCR) Load() RCR_Bits              { return RCR_Bits(r.U16.Load()) }
func (r *RCR) Store(b RCR_Bits)            { r.U16.Store(uint16(b)) }

type RCR_Mask struct{ mmio.UM16 }

func (rm RCR_Mask) Load() RCR_Bits   { return RCR_Bits(rm.UM16.Load()) }
func (rm RCR_Mask) Store(b RCR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) REP() RCR_Mask {
	return RCR_Mask{mmio.UM16{&p.RCR.U16, uint16(REP)}}
}

type CCR1_Bits uint16

func (b CCR1_Bits) Field(mask CCR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR1_Bits) J(v int) CCR1_Bits {
	return CCR1_Bits(bits.Make32(v, uint32(mask)))
}

type CCR1 struct{ mmio.U16 }

func (r *CCR1) Bits(mask CCR1_Bits) CCR1_Bits { return CCR1_Bits(r.U16.Bits(uint16(mask))) }
func (r *CCR1) StoreBits(mask, b CCR1_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CCR1) SetBits(mask CCR1_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *CCR1) ClearBits(mask CCR1_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *CCR1) Load() CCR1_Bits               { return CCR1_Bits(r.U16.Load()) }
func (r *CCR1) Store(b CCR1_Bits)             { r.U16.Store(uint16(b)) }

type CCR1_Mask struct{ mmio.UM16 }

func (rm CCR1_Mask) Load() CCR1_Bits   { return CCR1_Bits(rm.UM16.Load()) }
func (rm CCR1_Mask) Store(b CCR1_Bits) { rm.UM16.Store(uint16(b)) }

type CCR2_Bits uint16

func (b CCR2_Bits) Field(mask CCR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR2_Bits) J(v int) CCR2_Bits {
	return CCR2_Bits(bits.Make32(v, uint32(mask)))
}

type CCR2 struct{ mmio.U16 }

func (r *CCR2) Bits(mask CCR2_Bits) CCR2_Bits { return CCR2_Bits(r.U16.Bits(uint16(mask))) }
func (r *CCR2) StoreBits(mask, b CCR2_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CCR2) SetBits(mask CCR2_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *CCR2) ClearBits(mask CCR2_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *CCR2) Load() CCR2_Bits               { return CCR2_Bits(r.U16.Load()) }
func (r *CCR2) Store(b CCR2_Bits)             { r.U16.Store(uint16(b)) }

type CCR2_Mask struct{ mmio.UM16 }

func (rm CCR2_Mask) Load() CCR2_Bits   { return CCR2_Bits(rm.UM16.Load()) }
func (rm CCR2_Mask) Store(b CCR2_Bits) { rm.UM16.Store(uint16(b)) }

type CCR3_Bits uint16

func (b CCR3_Bits) Field(mask CCR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR3_Bits) J(v int) CCR3_Bits {
	return CCR3_Bits(bits.Make32(v, uint32(mask)))
}

type CCR3 struct{ mmio.U16 }

func (r *CCR3) Bits(mask CCR3_Bits) CCR3_Bits { return CCR3_Bits(r.U16.Bits(uint16(mask))) }
func (r *CCR3) StoreBits(mask, b CCR3_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CCR3) SetBits(mask CCR3_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *CCR3) ClearBits(mask CCR3_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *CCR3) Load() CCR3_Bits               { return CCR3_Bits(r.U16.Load()) }
func (r *CCR3) Store(b CCR3_Bits)             { r.U16.Store(uint16(b)) }

type CCR3_Mask struct{ mmio.UM16 }

func (rm CCR3_Mask) Load() CCR3_Bits   { return CCR3_Bits(rm.UM16.Load()) }
func (rm CCR3_Mask) Store(b CCR3_Bits) { rm.UM16.Store(uint16(b)) }

type CCR4_Bits uint16

func (b CCR4_Bits) Field(mask CCR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR4_Bits) J(v int) CCR4_Bits {
	return CCR4_Bits(bits.Make32(v, uint32(mask)))
}

type CCR4 struct{ mmio.U16 }

func (r *CCR4) Bits(mask CCR4_Bits) CCR4_Bits { return CCR4_Bits(r.U16.Bits(uint16(mask))) }
func (r *CCR4) StoreBits(mask, b CCR4_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CCR4) SetBits(mask CCR4_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *CCR4) ClearBits(mask CCR4_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *CCR4) Load() CCR4_Bits               { return CCR4_Bits(r.U16.Load()) }
func (r *CCR4) Store(b CCR4_Bits)             { r.U16.Store(uint16(b)) }

type CCR4_Mask struct{ mmio.UM16 }

func (rm CCR4_Mask) Load() CCR4_Bits   { return CCR4_Bits(rm.UM16.Load()) }
func (rm CCR4_Mask) Store(b CCR4_Bits) { rm.UM16.Store(uint16(b)) }

type BDTR_Bits uint16

func (b BDTR_Bits) Field(mask BDTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BDTR_Bits) J(v int) BDTR_Bits {
	return BDTR_Bits(bits.Make32(v, uint32(mask)))
}

type BDTR struct{ mmio.U16 }

func (r *BDTR) Bits(mask BDTR_Bits) BDTR_Bits { return BDTR_Bits(r.U16.Bits(uint16(mask))) }
func (r *BDTR) StoreBits(mask, b BDTR_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *BDTR) SetBits(mask BDTR_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *BDTR) ClearBits(mask BDTR_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *BDTR) Load() BDTR_Bits               { return BDTR_Bits(r.U16.Load()) }
func (r *BDTR) Store(b BDTR_Bits)             { r.U16.Store(uint16(b)) }

type BDTR_Mask struct{ mmio.UM16 }

func (rm BDTR_Mask) Load() BDTR_Bits   { return BDTR_Bits(rm.UM16.Load()) }
func (rm BDTR_Mask) Store(b BDTR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) DTG() BDTR_Mask {
	return BDTR_Mask{mmio.UM16{&p.BDTR.U16, uint16(DTG)}}
}

func (p *TIM_Periph) LOCK() BDTR_Mask {
	return BDTR_Mask{mmio.UM16{&p.BDTR.U16, uint16(LOCK)}}
}

func (p *TIM_Periph) OSSI() BDTR_Mask {
	return BDTR_Mask{mmio.UM16{&p.BDTR.U16, uint16(OSSI)}}
}

func (p *TIM_Periph) OSSR() BDTR_Mask {
	return BDTR_Mask{mmio.UM16{&p.BDTR.U16, uint16(OSSR)}}
}

func (p *TIM_Periph) BKE() BDTR_Mask {
	return BDTR_Mask{mmio.UM16{&p.BDTR.U16, uint16(BKE)}}
}

func (p *TIM_Periph) BKP() BDTR_Mask {
	return BDTR_Mask{mmio.UM16{&p.BDTR.U16, uint16(BKP)}}
}

func (p *TIM_Periph) AOE() BDTR_Mask {
	return BDTR_Mask{mmio.UM16{&p.BDTR.U16, uint16(AOE)}}
}

func (p *TIM_Periph) MOE() BDTR_Mask {
	return BDTR_Mask{mmio.UM16{&p.BDTR.U16, uint16(MOE)}}
}

type DCR_Bits uint16

func (b DCR_Bits) Field(mask DCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCR_Bits) J(v int) DCR_Bits {
	return DCR_Bits(bits.Make32(v, uint32(mask)))
}

type DCR struct{ mmio.U16 }

func (r *DCR) Bits(mask DCR_Bits) DCR_Bits { return DCR_Bits(r.U16.Bits(uint16(mask))) }
func (r *DCR) StoreBits(mask, b DCR_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DCR) SetBits(mask DCR_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *DCR) ClearBits(mask DCR_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *DCR) Load() DCR_Bits              { return DCR_Bits(r.U16.Load()) }
func (r *DCR) Store(b DCR_Bits)            { r.U16.Store(uint16(b)) }

type DCR_Mask struct{ mmio.UM16 }

func (rm DCR_Mask) Load() DCR_Bits   { return DCR_Bits(rm.UM16.Load()) }
func (rm DCR_Mask) Store(b DCR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) DBA() DCR_Mask {
	return DCR_Mask{mmio.UM16{&p.DCR.U16, uint16(DBA)}}
}

func (p *TIM_Periph) DBL() DCR_Mask {
	return DCR_Mask{mmio.UM16{&p.DCR.U16, uint16(DBL)}}
}

type DMAR_Bits uint16

func (b DMAR_Bits) Field(mask DMAR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DMAR_Bits) J(v int) DMAR_Bits {
	return DMAR_Bits(bits.Make32(v, uint32(mask)))
}

type DMAR struct{ mmio.U16 }

func (r *DMAR) Bits(mask DMAR_Bits) DMAR_Bits { return DMAR_Bits(r.U16.Bits(uint16(mask))) }
func (r *DMAR) StoreBits(mask, b DMAR_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DMAR) SetBits(mask DMAR_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *DMAR) ClearBits(mask DMAR_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *DMAR) Load() DMAR_Bits               { return DMAR_Bits(r.U16.Load()) }
func (r *DMAR) Store(b DMAR_Bits)             { r.U16.Store(uint16(b)) }

type DMAR_Mask struct{ mmio.UM16 }

func (rm DMAR_Mask) Load() DMAR_Bits   { return DMAR_Bits(rm.UM16.Load()) }
func (rm DMAR_Mask) Store(b DMAR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *TIM_Periph) DMAB() DMAR_Mask {
	return DMAR_Mask{mmio.UM16{&p.DMAR.U16, uint16(DMAB)}}
}
