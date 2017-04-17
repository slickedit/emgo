// +build f303xe

package rcc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type RCC_Periph struct {
	CR       CR
	CFGR     CFGR
	CIR      CIR
	APB2RSTR APB2RSTR
	APB1RSTR APB1RSTR
	AHBENR   AHBENR
	APB2ENR  APB2ENR
	APB1ENR  APB1ENR
	BDCR     BDCR
	CSR      CSR
	AHBRSTR  AHBRSTR
	CFGR2    CFGR2
	CFGR3    CFGR3
}

func (p *RCC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var RCC = (*RCC_Periph)(unsafe.Pointer(uintptr(mmap.RCC_BASE)))

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) HSION() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSION)}}
}

func (p *RCC_Periph) HSIRDY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSIRDY)}}
}

func (p *RCC_Periph) HSITRIM() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSITRIM)}}
}

func (p *RCC_Periph) HSICAL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSICAL)}}
}

func (p *RCC_Periph) HSEON() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSEON)}}
}

func (p *RCC_Periph) HSERDY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSERDY)}}
}

func (p *RCC_Periph) HSEBYP() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(HSEBYP)}}
}

func (p *RCC_Periph) CSSON() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CSSON)}}
}

func (p *RCC_Periph) PLLON() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PLLON)}}
}

func (p *RCC_Periph) PLLRDY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PLLRDY)}}
}

type CFGR_Bits uint32

func (b CFGR_Bits) Field(mask CFGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR_Bits) J(v int) CFGR_Bits {
	return CFGR_Bits(bits.Make32(v, uint32(mask)))
}

type CFGR struct{ mmio.U32 }

func (r *CFGR) Bits(mask CFGR_Bits) CFGR_Bits { return CFGR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFGR) StoreBits(mask, b CFGR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFGR) SetBits(mask CFGR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CFGR) ClearBits(mask CFGR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CFGR) Load() CFGR_Bits               { return CFGR_Bits(r.U32.Load()) }
func (r *CFGR) Store(b CFGR_Bits)             { r.U32.Store(uint32(b)) }

type CFGR_Mask struct{ mmio.UM32 }

func (rm CFGR_Mask) Load() CFGR_Bits   { return CFGR_Bits(rm.UM32.Load()) }
func (rm CFGR_Mask) Store(b CFGR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SW() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(SW)}}
}

func (p *RCC_Periph) SWS() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(SWS)}}
}

func (p *RCC_Periph) HPRE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(HPRE)}}
}

func (p *RCC_Periph) PPRE1() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PPRE1)}}
}

func (p *RCC_Periph) PPRE2() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PPRE2)}}
}

func (p *RCC_Periph) PLLSRC() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PLLSRC)}}
}

func (p *RCC_Periph) PLLXTPRE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PLLXTPRE)}}
}

func (p *RCC_Periph) PLLMULL() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PLLMULL)}}
}

func (p *RCC_Periph) PLLSRC_HSI_PREDIV() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PLLSRC_HSI_PREDIV)}}
}

func (p *RCC_Periph) USBPRE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(USBPRE)}}
}

func (p *RCC_Periph) I2SSRC() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(I2SSRC)}}
}

func (p *RCC_Periph) MCO() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(MCO)}}
}

func (p *RCC_Periph) MCOF() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(MCOF)}}
}

func (p *RCC_Periph) MCO_PRE_4() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(MCO_PRE_4)}}
}

func (p *RCC_Periph) MCO_PRE_16() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(MCO_PRE_16)}}
}

func (p *RCC_Periph) PLLNODIV() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PLLNODIV)}}
}

type CIR_Bits uint32

func (b CIR_Bits) Field(mask CIR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CIR_Bits) J(v int) CIR_Bits {
	return CIR_Bits(bits.Make32(v, uint32(mask)))
}

type CIR struct{ mmio.U32 }

func (r *CIR) Bits(mask CIR_Bits) CIR_Bits { return CIR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CIR) StoreBits(mask, b CIR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CIR) SetBits(mask CIR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CIR) ClearBits(mask CIR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CIR) Load() CIR_Bits              { return CIR_Bits(r.U32.Load()) }
func (r *CIR) Store(b CIR_Bits)            { r.U32.Store(uint32(b)) }

type CIR_Mask struct{ mmio.UM32 }

func (rm CIR_Mask) Load() CIR_Bits   { return CIR_Bits(rm.UM32.Load()) }
func (rm CIR_Mask) Store(b CIR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSIRDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSIRDYF)}}
}

func (p *RCC_Periph) LSERDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSERDYF)}}
}

func (p *RCC_Periph) HSIRDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSIRDYF)}}
}

func (p *RCC_Periph) HSERDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSERDYF)}}
}

func (p *RCC_Periph) PLLRDYF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(PLLRDYF)}}
}

func (p *RCC_Periph) CSSF() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(CSSF)}}
}

func (p *RCC_Periph) LSIRDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSIRDYIE)}}
}

func (p *RCC_Periph) LSERDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSERDYIE)}}
}

func (p *RCC_Periph) HSIRDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSIRDYIE)}}
}

func (p *RCC_Periph) HSERDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSERDYIE)}}
}

func (p *RCC_Periph) PLLRDYIE() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(PLLRDYIE)}}
}

func (p *RCC_Periph) LSIRDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSIRDYC)}}
}

func (p *RCC_Periph) LSERDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(LSERDYC)}}
}

func (p *RCC_Periph) HSIRDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSIRDYC)}}
}

func (p *RCC_Periph) HSERDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(HSERDYC)}}
}

func (p *RCC_Periph) PLLRDYC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(PLLRDYC)}}
}

func (p *RCC_Periph) CSSC() CIR_Mask {
	return CIR_Mask{mmio.UM32{&p.CIR.U32, uint32(CSSC)}}
}

type APB2RSTR_Bits uint32

func (b APB2RSTR_Bits) Field(mask APB2RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2RSTR_Bits) J(v int) APB2RSTR_Bits {
	return APB2RSTR_Bits(bits.Make32(v, uint32(mask)))
}

type APB2RSTR struct{ mmio.U32 }

func (r *APB2RSTR) Bits(mask APB2RSTR_Bits) APB2RSTR_Bits {
	return APB2RSTR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *APB2RSTR) StoreBits(mask, b APB2RSTR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB2RSTR) SetBits(mask APB2RSTR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *APB2RSTR) ClearBits(mask APB2RSTR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *APB2RSTR) Load() APB2RSTR_Bits             { return APB2RSTR_Bits(r.U32.Load()) }
func (r *APB2RSTR) Store(b APB2RSTR_Bits)           { r.U32.Store(uint32(b)) }

type APB2RSTR_Mask struct{ mmio.UM32 }

func (rm APB2RSTR_Mask) Load() APB2RSTR_Bits   { return APB2RSTR_Bits(rm.UM32.Load()) }
func (rm APB2RSTR_Mask) Store(b APB2RSTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SYSCFGRST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(SYSCFGRST)}}
}

func (p *RCC_Periph) TIM1RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM1RST)}}
}

func (p *RCC_Periph) SPI1RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(SPI1RST)}}
}

func (p *RCC_Periph) TIM8RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM8RST)}}
}

func (p *RCC_Periph) USART1RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(USART1RST)}}
}

func (p *RCC_Periph) SPI4RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(SPI4RST)}}
}

func (p *RCC_Periph) TIM15RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM15RST)}}
}

func (p *RCC_Periph) TIM16RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM16RST)}}
}

func (p *RCC_Periph) TIM17RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM17RST)}}
}

func (p *RCC_Periph) TIM20RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM20RST)}}
}

func (p *RCC_Periph) HRTIM1RST() APB2RSTR_Mask {
	return APB2RSTR_Mask{mmio.UM32{&p.APB2RSTR.U32, uint32(HRTIM1RST)}}
}

type APB1RSTR_Bits uint32

func (b APB1RSTR_Bits) Field(mask APB1RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1RSTR_Bits) J(v int) APB1RSTR_Bits {
	return APB1RSTR_Bits(bits.Make32(v, uint32(mask)))
}

type APB1RSTR struct{ mmio.U32 }

func (r *APB1RSTR) Bits(mask APB1RSTR_Bits) APB1RSTR_Bits {
	return APB1RSTR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *APB1RSTR) StoreBits(mask, b APB1RSTR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB1RSTR) SetBits(mask APB1RSTR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *APB1RSTR) ClearBits(mask APB1RSTR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *APB1RSTR) Load() APB1RSTR_Bits             { return APB1RSTR_Bits(r.U32.Load()) }
func (r *APB1RSTR) Store(b APB1RSTR_Bits)           { r.U32.Store(uint32(b)) }

type APB1RSTR_Mask struct{ mmio.UM32 }

func (rm APB1RSTR_Mask) Load() APB1RSTR_Bits   { return APB1RSTR_Bits(rm.UM32.Load()) }
func (rm APB1RSTR_Mask) Store(b APB1RSTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM2RST)}}
}

func (p *RCC_Periph) TIM3RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM3RST)}}
}

func (p *RCC_Periph) TIM4RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM4RST)}}
}

func (p *RCC_Periph) TIM6RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM6RST)}}
}

func (p *RCC_Periph) TIM7RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM7RST)}}
}

func (p *RCC_Periph) WWDGRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(WWDGRST)}}
}

func (p *RCC_Periph) SPI2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI2RST)}}
}

func (p *RCC_Periph) SPI3RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI3RST)}}
}

func (p *RCC_Periph) USART2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(USART2RST)}}
}

func (p *RCC_Periph) USART3RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(USART3RST)}}
}

func (p *RCC_Periph) UART4RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(UART4RST)}}
}

func (p *RCC_Periph) UART5RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(UART5RST)}}
}

func (p *RCC_Periph) I2C1RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C1RST)}}
}

func (p *RCC_Periph) I2C2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C2RST)}}
}

func (p *RCC_Periph) USBRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(USBRST)}}
}

func (p *RCC_Periph) CAN1RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(CAN1RST)}}
}

func (p *RCC_Periph) PWRRST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(PWRRST)}}
}

func (p *RCC_Periph) DAC1RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(DAC1RST)}}
}

func (p *RCC_Periph) I2C3RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C3RST)}}
}

func (p *RCC_Periph) DAC2RST() APB1RSTR_Mask {
	return APB1RSTR_Mask{mmio.UM32{&p.APB1RSTR.U32, uint32(DAC2RST)}}
}

type AHBENR_Bits uint32

func (b AHBENR_Bits) Field(mask AHBENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBENR_Bits) J(v int) AHBENR_Bits {
	return AHBENR_Bits(bits.Make32(v, uint32(mask)))
}

type AHBENR struct{ mmio.U32 }

func (r *AHBENR) Bits(mask AHBENR_Bits) AHBENR_Bits { return AHBENR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AHBENR) StoreBits(mask, b AHBENR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AHBENR) SetBits(mask AHBENR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *AHBENR) ClearBits(mask AHBENR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *AHBENR) Load() AHBENR_Bits                 { return AHBENR_Bits(r.U32.Load()) }
func (r *AHBENR) Store(b AHBENR_Bits)               { r.U32.Store(uint32(b)) }

type AHBENR_Mask struct{ mmio.UM32 }

func (rm AHBENR_Mask) Load() AHBENR_Bits   { return AHBENR_Bits(rm.UM32.Load()) }
func (rm AHBENR_Mask) Store(b AHBENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) DMA1EN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(DMA1EN)}}
}

func (p *RCC_Periph) DMA2EN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(DMA2EN)}}
}

func (p *RCC_Periph) SRAMEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(SRAMEN)}}
}

func (p *RCC_Periph) FLITFEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(FLITFEN)}}
}

func (p *RCC_Periph) FMCEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(FMCEN)}}
}

func (p *RCC_Periph) CRCEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(CRCEN)}}
}

func (p *RCC_Periph) GPIOHEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOHEN)}}
}

func (p *RCC_Periph) GPIOAEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOAEN)}}
}

func (p *RCC_Periph) GPIOBEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOBEN)}}
}

func (p *RCC_Periph) GPIOCEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOCEN)}}
}

func (p *RCC_Periph) GPIODEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIODEN)}}
}

func (p *RCC_Periph) GPIOEEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOEEN)}}
}

func (p *RCC_Periph) GPIOFEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOFEN)}}
}

func (p *RCC_Periph) GPIOGEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(GPIOGEN)}}
}

func (p *RCC_Periph) TSEN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(TSEN)}}
}

func (p *RCC_Periph) ADC12EN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(ADC12EN)}}
}

func (p *RCC_Periph) ADC34EN() AHBENR_Mask {
	return AHBENR_Mask{mmio.UM32{&p.AHBENR.U32, uint32(ADC34EN)}}
}

type APB2ENR_Bits uint32

func (b APB2ENR_Bits) Field(mask APB2ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2ENR_Bits) J(v int) APB2ENR_Bits {
	return APB2ENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB2ENR struct{ mmio.U32 }

func (r *APB2ENR) Bits(mask APB2ENR_Bits) APB2ENR_Bits { return APB2ENR_Bits(r.U32.Bits(uint32(mask))) }
func (r *APB2ENR) StoreBits(mask, b APB2ENR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB2ENR) SetBits(mask APB2ENR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *APB2ENR) ClearBits(mask APB2ENR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *APB2ENR) Load() APB2ENR_Bits                  { return APB2ENR_Bits(r.U32.Load()) }
func (r *APB2ENR) Store(b APB2ENR_Bits)                { r.U32.Store(uint32(b)) }

type APB2ENR_Mask struct{ mmio.UM32 }

func (rm APB2ENR_Mask) Load() APB2ENR_Bits   { return APB2ENR_Bits(rm.UM32.Load()) }
func (rm APB2ENR_Mask) Store(b APB2ENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SYSCFGEN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(SYSCFGEN)}}
}

func (p *RCC_Periph) TIM1EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(TIM1EN)}}
}

func (p *RCC_Periph) SPI1EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(SPI1EN)}}
}

func (p *RCC_Periph) TIM8EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(TIM8EN)}}
}

func (p *RCC_Periph) USART1EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(USART1EN)}}
}

func (p *RCC_Periph) SPI4EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(SPI4EN)}}
}

func (p *RCC_Periph) TIM15EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(TIM15EN)}}
}

func (p *RCC_Periph) TIM16EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(TIM16EN)}}
}

func (p *RCC_Periph) TIM17EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(TIM17EN)}}
}

func (p *RCC_Periph) TIM20EN() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(TIM20EN)}}
}

func (p *RCC_Periph) HRTIM1() APB2ENR_Mask {
	return APB2ENR_Mask{mmio.UM32{&p.APB2ENR.U32, uint32(HRTIM1)}}
}

type APB1ENR_Bits uint32

func (b APB1ENR_Bits) Field(mask APB1ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1ENR_Bits) J(v int) APB1ENR_Bits {
	return APB1ENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB1ENR struct{ mmio.U32 }

func (r *APB1ENR) Bits(mask APB1ENR_Bits) APB1ENR_Bits { return APB1ENR_Bits(r.U32.Bits(uint32(mask))) }
func (r *APB1ENR) StoreBits(mask, b APB1ENR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *APB1ENR) SetBits(mask APB1ENR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *APB1ENR) ClearBits(mask APB1ENR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *APB1ENR) Load() APB1ENR_Bits                  { return APB1ENR_Bits(r.U32.Load()) }
func (r *APB1ENR) Store(b APB1ENR_Bits)                { r.U32.Store(uint32(b)) }

type APB1ENR_Mask struct{ mmio.UM32 }

func (rm APB1ENR_Mask) Load() APB1ENR_Bits   { return APB1ENR_Bits(rm.UM32.Load()) }
func (rm APB1ENR_Mask) Store(b APB1ENR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM2EN)}}
}

func (p *RCC_Periph) TIM3EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM3EN)}}
}

func (p *RCC_Periph) TIM4EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM4EN)}}
}

func (p *RCC_Periph) TIM6EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM6EN)}}
}

func (p *RCC_Periph) TIM7EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(TIM7EN)}}
}

func (p *RCC_Periph) WWDGEN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(WWDGEN)}}
}

func (p *RCC_Periph) SPI2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(SPI2EN)}}
}

func (p *RCC_Periph) SPI3EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(SPI3EN)}}
}

func (p *RCC_Periph) USART2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(USART2EN)}}
}

func (p *RCC_Periph) USART3EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(USART3EN)}}
}

func (p *RCC_Periph) UART4EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(UART4EN)}}
}

func (p *RCC_Periph) UART5EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(UART5EN)}}
}

func (p *RCC_Periph) I2C1EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(I2C1EN)}}
}

func (p *RCC_Periph) I2C2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(I2C2EN)}}
}

func (p *RCC_Periph) USBEN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(USBEN)}}
}

func (p *RCC_Periph) CAN1EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(CAN1EN)}}
}

func (p *RCC_Periph) DAC2EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(DAC2EN)}}
}

func (p *RCC_Periph) PWREN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(PWREN)}}
}

func (p *RCC_Periph) DAC1EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(DAC1EN)}}
}

func (p *RCC_Periph) I2C3EN() APB1ENR_Mask {
	return APB1ENR_Mask{mmio.UM32{&p.APB1ENR.U32, uint32(I2C3EN)}}
}

type BDCR_Bits uint32

func (b BDCR_Bits) Field(mask BDCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BDCR_Bits) J(v int) BDCR_Bits {
	return BDCR_Bits(bits.Make32(v, uint32(mask)))
}

type BDCR struct{ mmio.U32 }

func (r *BDCR) Bits(mask BDCR_Bits) BDCR_Bits { return BDCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BDCR) StoreBits(mask, b BDCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BDCR) SetBits(mask BDCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BDCR) ClearBits(mask BDCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BDCR) Load() BDCR_Bits               { return BDCR_Bits(r.U32.Load()) }
func (r *BDCR) Store(b BDCR_Bits)             { r.U32.Store(uint32(b)) }

type BDCR_Mask struct{ mmio.UM32 }

func (rm BDCR_Mask) Load() BDCR_Bits   { return BDCR_Bits(rm.UM32.Load()) }
func (rm BDCR_Mask) Store(b BDCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSEON() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(LSEON)}}
}

func (p *RCC_Periph) LSERDY() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(LSERDY)}}
}

func (p *RCC_Periph) LSEBYP() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(LSEBYP)}}
}

func (p *RCC_Periph) LSEDRV() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(LSEDRV)}}
}

func (p *RCC_Periph) RTCSEL() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(RTCSEL)}}
}

func (p *RCC_Periph) RTCEN() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(RTCEN)}}
}

func (p *RCC_Periph) BDRST() BDCR_Mask {
	return BDCR_Mask{mmio.UM32{&p.BDCR.U32, uint32(BDRST)}}
}

type CSR_Bits uint32

func (b CSR_Bits) Field(mask CSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR_Bits) J(v int) CSR_Bits {
	return CSR_Bits(bits.Make32(v, uint32(mask)))
}

type CSR struct{ mmio.U32 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.U32.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.U32.Store(uint32(b)) }

type CSR_Mask struct{ mmio.UM32 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM32.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSION() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LSION)}}
}

func (p *RCC_Periph) LSIRDY() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LSIRDY)}}
}

func (p *RCC_Periph) RMVF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(RMVF)}}
}

func (p *RCC_Periph) OBLRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OBLRSTF)}}
}

func (p *RCC_Periph) PINRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(PINRSTF)}}
}

func (p *RCC_Periph) PORRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(PORRSTF)}}
}

func (p *RCC_Periph) SFTRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(SFTRSTF)}}
}

func (p *RCC_Periph) IWDGRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(IWDGRSTF)}}
}

func (p *RCC_Periph) WWDGRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(WWDGRSTF)}}
}

func (p *RCC_Periph) LPWRRSTF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LPWRRSTF)}}
}

type AHBRSTR_Bits uint32

func (b AHBRSTR_Bits) Field(mask AHBRSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBRSTR_Bits) J(v int) AHBRSTR_Bits {
	return AHBRSTR_Bits(bits.Make32(v, uint32(mask)))
}

type AHBRSTR struct{ mmio.U32 }

func (r *AHBRSTR) Bits(mask AHBRSTR_Bits) AHBRSTR_Bits { return AHBRSTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AHBRSTR) StoreBits(mask, b AHBRSTR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AHBRSTR) SetBits(mask AHBRSTR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *AHBRSTR) ClearBits(mask AHBRSTR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *AHBRSTR) Load() AHBRSTR_Bits                  { return AHBRSTR_Bits(r.U32.Load()) }
func (r *AHBRSTR) Store(b AHBRSTR_Bits)                { r.U32.Store(uint32(b)) }

type AHBRSTR_Mask struct{ mmio.UM32 }

func (rm AHBRSTR_Mask) Load() AHBRSTR_Bits   { return AHBRSTR_Bits(rm.UM32.Load()) }
func (rm AHBRSTR_Mask) Store(b AHBRSTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) FMCRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(FMCRST)}}
}

func (p *RCC_Periph) GPIOHRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOHRST)}}
}

func (p *RCC_Periph) GPIOARST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOARST)}}
}

func (p *RCC_Periph) GPIOBRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOBRST)}}
}

func (p *RCC_Periph) GPIOCRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOCRST)}}
}

func (p *RCC_Periph) GPIOERST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOERST)}}
}

func (p *RCC_Periph) GPIOFRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOFRST)}}
}

func (p *RCC_Periph) GPIOGRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOGRST)}}
}

func (p *RCC_Periph) TSRST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(TSRST)}}
}

func (p *RCC_Periph) ADC12RST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(ADC12RST)}}
}

func (p *RCC_Periph) ADC34RST() AHBRSTR_Mask {
	return AHBRSTR_Mask{mmio.UM32{&p.AHBRSTR.U32, uint32(ADC34RST)}}
}

type CFGR2_Bits uint32

func (b CFGR2_Bits) Field(mask CFGR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR2_Bits) J(v int) CFGR2_Bits {
	return CFGR2_Bits(bits.Make32(v, uint32(mask)))
}

type CFGR2 struct{ mmio.U32 }

func (r *CFGR2) Bits(mask CFGR2_Bits) CFGR2_Bits { return CFGR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFGR2) StoreBits(mask, b CFGR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFGR2) SetBits(mask CFGR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CFGR2) ClearBits(mask CFGR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CFGR2) Load() CFGR2_Bits                { return CFGR2_Bits(r.U32.Load()) }
func (r *CFGR2) Store(b CFGR2_Bits)              { r.U32.Store(uint32(b)) }

type CFGR2_Mask struct{ mmio.UM32 }

func (rm CFGR2_Mask) Load() CFGR2_Bits   { return CFGR2_Bits(rm.UM32.Load()) }
func (rm CFGR2_Mask) Store(b CFGR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) PREDIV1() CFGR2_Mask {
	return CFGR2_Mask{mmio.UM32{&p.CFGR2.U32, uint32(PREDIV1)}}
}

func (p *RCC_Periph) ADCPRE12() CFGR2_Mask {
	return CFGR2_Mask{mmio.UM32{&p.CFGR2.U32, uint32(ADCPRE12)}}
}

func (p *RCC_Periph) ADCPRE34() CFGR2_Mask {
	return CFGR2_Mask{mmio.UM32{&p.CFGR2.U32, uint32(ADCPRE34)}}
}

type CFGR3_Bits uint32

func (b CFGR3_Bits) Field(mask CFGR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR3_Bits) J(v int) CFGR3_Bits {
	return CFGR3_Bits(bits.Make32(v, uint32(mask)))
}

type CFGR3 struct{ mmio.U32 }

func (r *CFGR3) Bits(mask CFGR3_Bits) CFGR3_Bits { return CFGR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFGR3) StoreBits(mask, b CFGR3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFGR3) SetBits(mask CFGR3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CFGR3) ClearBits(mask CFGR3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CFGR3) Load() CFGR3_Bits                { return CFGR3_Bits(r.U32.Load()) }
func (r *CFGR3) Store(b CFGR3_Bits)              { r.U32.Store(uint32(b)) }

type CFGR3_Mask struct{ mmio.UM32 }

func (rm CFGR3_Mask) Load() CFGR3_Bits   { return CFGR3_Bits(rm.UM32.Load()) }
func (rm CFGR3_Mask) Store(b CFGR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) USART1SW() CFGR3_Mask {
	return CFGR3_Mask{mmio.UM32{&p.CFGR3.U32, uint32(USART1SW)}}
}

func (p *RCC_Periph) I2CSW() CFGR3_Mask {
	return CFGR3_Mask{mmio.UM32{&p.CFGR3.U32, uint32(I2CSW)}}
}

func (p *RCC_Periph) TIMSW() CFGR3_Mask {
	return CFGR3_Mask{mmio.UM32{&p.CFGR3.U32, uint32(TIMSW)}}
}

func (p *RCC_Periph) TIM20SW() CFGR3_Mask {
	return CFGR3_Mask{mmio.UM32{&p.CFGR3.U32, uint32(TIM20SW)}}
}

func (p *RCC_Periph) TIM2SW() CFGR3_Mask {
	return CFGR3_Mask{mmio.UM32{&p.CFGR3.U32, uint32(TIM2SW)}}
}

func (p *RCC_Periph) TIM3SW() CFGR3_Mask {
	return CFGR3_Mask{mmio.UM32{&p.CFGR3.U32, uint32(TIM3SW)}}
}

func (p *RCC_Periph) HRTIM1SW() CFGR3_Mask {
	return CFGR3_Mask{mmio.UM32{&p.CFGR3.U32, uint32(HRTIM1SW)}}
}

func (p *RCC_Periph) USART2SW() CFGR3_Mask {
	return CFGR3_Mask{mmio.UM32{&p.CFGR3.U32, uint32(USART2SW)}}
}

func (p *RCC_Periph) USART3SW() CFGR3_Mask {
	return CFGR3_Mask{mmio.UM32{&p.CFGR3.U32, uint32(USART3SW)}}
}

func (p *RCC_Periph) UART4SW() CFGR3_Mask {
	return CFGR3_Mask{mmio.UM32{&p.CFGR3.U32, uint32(UART4SW)}}
}

func (p *RCC_Periph) UART5SW() CFGR3_Mask {
	return CFGR3_Mask{mmio.UM32{&p.CFGR3.U32, uint32(UART5SW)}}
}
