// +build l1xx_md

package syscfg

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type SYSCFG_Periph struct {
	MEMRMP MEMRMP
	PMC    PMC
	EXTICR [4]EXTICR
}

func (p *SYSCFG_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SYSCFG = (*SYSCFG_Periph)(unsafe.Pointer(uintptr(mmap.SYSCFG_BASE)))

type MEMRMP_Bits uint32

func (b MEMRMP_Bits) Field(mask MEMRMP_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MEMRMP_Bits) J(v int) MEMRMP_Bits {
	return MEMRMP_Bits(bits.Make32(v, uint32(mask)))
}

type MEMRMP struct{ mmio.U32 }

func (r *MEMRMP) Bits(mask MEMRMP_Bits) MEMRMP_Bits { return MEMRMP_Bits(r.U32.Bits(uint32(mask))) }
func (r *MEMRMP) StoreBits(mask, b MEMRMP_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MEMRMP) SetBits(mask MEMRMP_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *MEMRMP) ClearBits(mask MEMRMP_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *MEMRMP) Load() MEMRMP_Bits                 { return MEMRMP_Bits(r.U32.Load()) }
func (r *MEMRMP) Store(b MEMRMP_Bits)               { r.U32.Store(uint32(b)) }

type MEMRMP_Mask struct{ mmio.UM32 }

func (rm MEMRMP_Mask) Load() MEMRMP_Bits   { return MEMRMP_Bits(rm.UM32.Load()) }
func (rm MEMRMP_Mask) Store(b MEMRMP_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) MEM_MODE() MEMRMP_Mask {
	return MEMRMP_Mask{mmio.UM32{&p.MEMRMP.U32, uint32(MEM_MODE)}}
}

func (p *SYSCFG_Periph) BOOT_MODE() MEMRMP_Mask {
	return MEMRMP_Mask{mmio.UM32{&p.MEMRMP.U32, uint32(BOOT_MODE)}}
}

type PMC_Bits uint32

func (b PMC_Bits) Field(mask PMC_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PMC_Bits) J(v int) PMC_Bits {
	return PMC_Bits(bits.Make32(v, uint32(mask)))
}

type PMC struct{ mmio.U32 }

func (r *PMC) Bits(mask PMC_Bits) PMC_Bits { return PMC_Bits(r.U32.Bits(uint32(mask))) }
func (r *PMC) StoreBits(mask, b PMC_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PMC) SetBits(mask PMC_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *PMC) ClearBits(mask PMC_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *PMC) Load() PMC_Bits              { return PMC_Bits(r.U32.Load()) }
func (r *PMC) Store(b PMC_Bits)            { r.U32.Store(uint32(b)) }

type PMC_Mask struct{ mmio.UM32 }

func (rm PMC_Mask) Load() PMC_Bits   { return PMC_Bits(rm.UM32.Load()) }
func (rm PMC_Mask) Store(b PMC_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) USB_PU() PMC_Mask {
	return PMC_Mask{mmio.UM32{&p.PMC.U32, uint32(USB_PU)}}
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

func (p *SYSCFG_Periph) EXTI0(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI0)}}
}

func (p *SYSCFG_Periph) EXTI1(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI1)}}
}

func (p *SYSCFG_Periph) EXTI2(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI2)}}
}

func (p *SYSCFG_Periph) EXTI3(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI3)}}
}
