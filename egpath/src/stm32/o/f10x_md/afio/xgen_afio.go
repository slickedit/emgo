package afio

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type AFIO_Periph struct {
	EVCR   EVCR
	MAPR   MAPR
	EXTICR [4]EXTICR
	_      uint32
	MAPR2  MAPR2
}

func (p *AFIO_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var AFIO = (*AFIO_Periph)(unsafe.Pointer(uintptr(mmap.AFIO_BASE)))

type EVCR_Bits uint32

func (b EVCR_Bits) Field(mask EVCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EVCR_Bits) J(v int) EVCR_Bits {
	return EVCR_Bits(bits.Make32(v, uint32(mask)))
}

type EVCR struct{ mmio.U32 }

func (r *EVCR) Bits(mask EVCR_Bits) EVCR_Bits { return EVCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *EVCR) StoreBits(mask, b EVCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *EVCR) SetBits(mask EVCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *EVCR) ClearBits(mask EVCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *EVCR) Load() EVCR_Bits               { return EVCR_Bits(r.U32.Load()) }
func (r *EVCR) Store(b EVCR_Bits)             { r.U32.Store(uint32(b)) }

type EVCR_Mask struct{ mmio.UM32 }

func (rm EVCR_Mask) Load() EVCR_Bits   { return EVCR_Bits(rm.UM32.Load()) }
func (rm EVCR_Mask) Store(b EVCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *AFIO_Periph) PIN() EVCR_Mask {
	return EVCR_Mask{mmio.UM32{&p.EVCR.U32, uint32(PIN)}}
}

func (p *AFIO_Periph) PORT() EVCR_Mask {
	return EVCR_Mask{mmio.UM32{&p.EVCR.U32, uint32(PORT)}}
}

func (p *AFIO_Periph) EVOE() EVCR_Mask {
	return EVCR_Mask{mmio.UM32{&p.EVCR.U32, uint32(EVOE)}}
}

type MAPR_Bits uint32

func (b MAPR_Bits) Field(mask MAPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MAPR_Bits) J(v int) MAPR_Bits {
	return MAPR_Bits(bits.Make32(v, uint32(mask)))
}

type MAPR struct{ mmio.U32 }

func (r *MAPR) Bits(mask MAPR_Bits) MAPR_Bits { return MAPR_Bits(r.U32.Bits(uint32(mask))) }
func (r *MAPR) StoreBits(mask, b MAPR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MAPR) SetBits(mask MAPR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *MAPR) ClearBits(mask MAPR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *MAPR) Load() MAPR_Bits               { return MAPR_Bits(r.U32.Load()) }
func (r *MAPR) Store(b MAPR_Bits)             { r.U32.Store(uint32(b)) }

type MAPR_Mask struct{ mmio.UM32 }

func (rm MAPR_Mask) Load() MAPR_Bits   { return MAPR_Bits(rm.UM32.Load()) }
func (rm MAPR_Mask) Store(b MAPR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *AFIO_Periph) SPI1_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(SPI1_REMAP)}}
}

func (p *AFIO_Periph) I2C1_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(I2C1_REMAP)}}
}

func (p *AFIO_Periph) USART1_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(USART1_REMAP)}}
}

func (p *AFIO_Periph) USART2_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(USART2_REMAP)}}
}

func (p *AFIO_Periph) USART3_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(USART3_REMAP)}}
}

func (p *AFIO_Periph) TIM1_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(TIM1_REMAP)}}
}

func (p *AFIO_Periph) TIM2_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(TIM2_REMAP)}}
}

func (p *AFIO_Periph) TIM3_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(TIM3_REMAP)}}
}

func (p *AFIO_Periph) TIM4_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(TIM4_REMAP)}}
}

func (p *AFIO_Periph) CAN_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(CAN_REMAP)}}
}

func (p *AFIO_Periph) PD01_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(PD01_REMAP)}}
}

func (p *AFIO_Periph) TIM5CH4_IREMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(TIM5CH4_IREMAP)}}
}

func (p *AFIO_Periph) ADC1_ETRGINJ_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(ADC1_ETRGINJ_REMAP)}}
}

func (p *AFIO_Periph) ADC1_ETRGREG_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(ADC1_ETRGREG_REMAP)}}
}

func (p *AFIO_Periph) ADC2_ETRGINJ_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(ADC2_ETRGINJ_REMAP)}}
}

func (p *AFIO_Periph) ADC2_ETRGREG_REMAP() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(ADC2_ETRGREG_REMAP)}}
}

func (p *AFIO_Periph) SWJ_CFG() MAPR_Mask {
	return MAPR_Mask{mmio.UM32{&p.MAPR.U32, uint32(SWJ_CFG)}}
}

type EXTICR_Bits uint32

func (b EXTICR_Bits) Field(mask EXTICR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EXTICR_Bits) J(v int) EXTICR_Bits {
	return EXTICR_Bits(bits.Make32(v, uint32(mask)))
}

type EXTICR struct{ mmio.U32 }

func (r *EXTICR) Bits(mask EXTICR_Bits) EXTICR_Bits { return EXTICR_Bits(r.U32.Bits(uint32(mask))) }
func (r *EXTICR) StoreBits(mask, b EXTICR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *EXTICR) SetBits(mask EXTICR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *EXTICR) ClearBits(mask EXTICR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *EXTICR) Load() EXTICR_Bits                 { return EXTICR_Bits(r.U32.Load()) }
func (r *EXTICR) Store(b EXTICR_Bits)               { r.U32.Store(uint32(b)) }

type EXTICR_Mask struct{ mmio.UM32 }

func (rm EXTICR_Mask) Load() EXTICR_Bits   { return EXTICR_Bits(rm.UM32.Load()) }
func (rm EXTICR_Mask) Store(b EXTICR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *AFIO_Periph) EXTI0(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI0)}}
}

func (p *AFIO_Periph) EXTI1(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI1)}}
}

func (p *AFIO_Periph) EXTI2(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI2)}}
}

func (p *AFIO_Periph) EXTI3(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI3)}}
}

type MAPR2_Bits uint32

func (b MAPR2_Bits) Field(mask MAPR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MAPR2_Bits) J(v int) MAPR2_Bits {
	return MAPR2_Bits(bits.Make32(v, uint32(mask)))
}

type MAPR2 struct{ mmio.U32 }

func (r *MAPR2) Bits(mask MAPR2_Bits) MAPR2_Bits { return MAPR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *MAPR2) StoreBits(mask, b MAPR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MAPR2) SetBits(mask MAPR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *MAPR2) ClearBits(mask MAPR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *MAPR2) Load() MAPR2_Bits                { return MAPR2_Bits(r.U32.Load()) }
func (r *MAPR2) Store(b MAPR2_Bits)              { r.U32.Store(uint32(b)) }

type MAPR2_Mask struct{ mmio.UM32 }

func (rm MAPR2_Mask) Load() MAPR2_Bits   { return MAPR2_Bits(rm.UM32.Load()) }
func (rm MAPR2_Mask) Store(b MAPR2_Bits) { rm.UM32.Store(uint32(b)) }
