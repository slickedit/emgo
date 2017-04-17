// +build f40_41xxx

package dma

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type DMA_Periph struct {
	LISR  LISR
	HISR  HISR
	LIFCR LIFCR
	HIFCR HIFCR
}

func (p *DMA_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DMA1 = (*DMA_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_BASE)))

//emgo:const
var DMA2 = (*DMA_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_BASE)))

type LISR_Bits uint32

func (b LISR_Bits) Field(mask LISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LISR_Bits) J(v int) LISR_Bits {
	return LISR_Bits(bits.Make32(v, uint32(mask)))
}

type LISR struct{ mmio.U32 }

func (r *LISR) Bits(mask LISR_Bits) LISR_Bits { return LISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LISR) StoreBits(mask, b LISR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LISR) SetBits(mask LISR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *LISR) ClearBits(mask LISR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *LISR) Load() LISR_Bits               { return LISR_Bits(r.U32.Load()) }
func (r *LISR) Store(b LISR_Bits)             { r.U32.Store(uint32(b)) }

type LISR_Mask struct{ mmio.UM32 }

func (rm LISR_Mask) Load() LISR_Bits   { return LISR_Bits(rm.UM32.Load()) }
func (rm LISR_Mask) Store(b LISR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) TCIF3() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(TCIF3)}}
}

func (p *DMA_Periph) HTIF3() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(HTIF3)}}
}

func (p *DMA_Periph) TEIF3() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(TEIF3)}}
}

func (p *DMA_Periph) DMEIF3() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(DMEIF3)}}
}

func (p *DMA_Periph) FEIF3() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(FEIF3)}}
}

func (p *DMA_Periph) TCIF2() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(TCIF2)}}
}

func (p *DMA_Periph) HTIF2() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(HTIF2)}}
}

func (p *DMA_Periph) TEIF2() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(TEIF2)}}
}

func (p *DMA_Periph) DMEIF2() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(DMEIF2)}}
}

func (p *DMA_Periph) FEIF2() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(FEIF2)}}
}

func (p *DMA_Periph) TCIF1() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(TCIF1)}}
}

func (p *DMA_Periph) HTIF1() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(HTIF1)}}
}

func (p *DMA_Periph) TEIF1() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(TEIF1)}}
}

func (p *DMA_Periph) DMEIF1() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(DMEIF1)}}
}

func (p *DMA_Periph) FEIF1() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(FEIF1)}}
}

func (p *DMA_Periph) TCIF0() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(TCIF0)}}
}

func (p *DMA_Periph) HTIF0() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(HTIF0)}}
}

func (p *DMA_Periph) TEIF0() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(TEIF0)}}
}

func (p *DMA_Periph) DMEIF0() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(DMEIF0)}}
}

func (p *DMA_Periph) FEIF0() LISR_Mask {
	return LISR_Mask{mmio.UM32{&p.LISR.U32, uint32(FEIF0)}}
}

type HISR_Bits uint32

func (b HISR_Bits) Field(mask HISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HISR_Bits) J(v int) HISR_Bits {
	return HISR_Bits(bits.Make32(v, uint32(mask)))
}

type HISR struct{ mmio.U32 }

func (r *HISR) Bits(mask HISR_Bits) HISR_Bits { return HISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *HISR) StoreBits(mask, b HISR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HISR) SetBits(mask HISR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *HISR) ClearBits(mask HISR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *HISR) Load() HISR_Bits               { return HISR_Bits(r.U32.Load()) }
func (r *HISR) Store(b HISR_Bits)             { r.U32.Store(uint32(b)) }

type HISR_Mask struct{ mmio.UM32 }

func (rm HISR_Mask) Load() HISR_Bits   { return HISR_Bits(rm.UM32.Load()) }
func (rm HISR_Mask) Store(b HISR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) TCIF7() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(TCIF7)}}
}

func (p *DMA_Periph) HTIF7() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(HTIF7)}}
}

func (p *DMA_Periph) TEIF7() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(TEIF7)}}
}

func (p *DMA_Periph) DMEIF7() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(DMEIF7)}}
}

func (p *DMA_Periph) FEIF7() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(FEIF7)}}
}

func (p *DMA_Periph) TCIF6() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(TCIF6)}}
}

func (p *DMA_Periph) HTIF6() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(HTIF6)}}
}

func (p *DMA_Periph) TEIF6() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(TEIF6)}}
}

func (p *DMA_Periph) DMEIF6() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(DMEIF6)}}
}

func (p *DMA_Periph) FEIF6() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(FEIF6)}}
}

func (p *DMA_Periph) TCIF5() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(TCIF5)}}
}

func (p *DMA_Periph) HTIF5() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(HTIF5)}}
}

func (p *DMA_Periph) TEIF5() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(TEIF5)}}
}

func (p *DMA_Periph) DMEIF5() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(DMEIF5)}}
}

func (p *DMA_Periph) FEIF5() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(FEIF5)}}
}

func (p *DMA_Periph) TCIF4() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(TCIF4)}}
}

func (p *DMA_Periph) HTIF4() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(HTIF4)}}
}

func (p *DMA_Periph) TEIF4() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(TEIF4)}}
}

func (p *DMA_Periph) DMEIF4() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(DMEIF4)}}
}

func (p *DMA_Periph) FEIF4() HISR_Mask {
	return HISR_Mask{mmio.UM32{&p.HISR.U32, uint32(FEIF4)}}
}

type LIFCR_Bits uint32

func (b LIFCR_Bits) Field(mask LIFCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LIFCR_Bits) J(v int) LIFCR_Bits {
	return LIFCR_Bits(bits.Make32(v, uint32(mask)))
}

type LIFCR struct{ mmio.U32 }

func (r *LIFCR) Bits(mask LIFCR_Bits) LIFCR_Bits { return LIFCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LIFCR) StoreBits(mask, b LIFCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LIFCR) SetBits(mask LIFCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *LIFCR) ClearBits(mask LIFCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *LIFCR) Load() LIFCR_Bits                { return LIFCR_Bits(r.U32.Load()) }
func (r *LIFCR) Store(b LIFCR_Bits)              { r.U32.Store(uint32(b)) }

type LIFCR_Mask struct{ mmio.UM32 }

func (rm LIFCR_Mask) Load() LIFCR_Bits   { return LIFCR_Bits(rm.UM32.Load()) }
func (rm LIFCR_Mask) Store(b LIFCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) CTCIF3() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CTCIF3)}}
}

func (p *DMA_Periph) CHTIF3() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CHTIF3)}}
}

func (p *DMA_Periph) CTEIF3() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CTEIF3)}}
}

func (p *DMA_Periph) CDMEIF3() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CDMEIF3)}}
}

func (p *DMA_Periph) CFEIF3() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CFEIF3)}}
}

func (p *DMA_Periph) CTCIF2() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CTCIF2)}}
}

func (p *DMA_Periph) CHTIF2() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CHTIF2)}}
}

func (p *DMA_Periph) CTEIF2() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CTEIF2)}}
}

func (p *DMA_Periph) CDMEIF2() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CDMEIF2)}}
}

func (p *DMA_Periph) CFEIF2() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CFEIF2)}}
}

func (p *DMA_Periph) CTCIF1() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CTCIF1)}}
}

func (p *DMA_Periph) CHTIF1() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CHTIF1)}}
}

func (p *DMA_Periph) CTEIF1() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CTEIF1)}}
}

func (p *DMA_Periph) CDMEIF1() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CDMEIF1)}}
}

func (p *DMA_Periph) CFEIF1() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CFEIF1)}}
}

func (p *DMA_Periph) CTCIF0() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CTCIF0)}}
}

func (p *DMA_Periph) CHTIF0() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CHTIF0)}}
}

func (p *DMA_Periph) CTEIF0() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CTEIF0)}}
}

func (p *DMA_Periph) CDMEIF0() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CDMEIF0)}}
}

func (p *DMA_Periph) CFEIF0() LIFCR_Mask {
	return LIFCR_Mask{mmio.UM32{&p.LIFCR.U32, uint32(CFEIF0)}}
}

type HIFCR_Bits uint32

func (b HIFCR_Bits) Field(mask HIFCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HIFCR_Bits) J(v int) HIFCR_Bits {
	return HIFCR_Bits(bits.Make32(v, uint32(mask)))
}

type HIFCR struct{ mmio.U32 }

func (r *HIFCR) Bits(mask HIFCR_Bits) HIFCR_Bits { return HIFCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *HIFCR) StoreBits(mask, b HIFCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HIFCR) SetBits(mask HIFCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *HIFCR) ClearBits(mask HIFCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *HIFCR) Load() HIFCR_Bits                { return HIFCR_Bits(r.U32.Load()) }
func (r *HIFCR) Store(b HIFCR_Bits)              { r.U32.Store(uint32(b)) }

type HIFCR_Mask struct{ mmio.UM32 }

func (rm HIFCR_Mask) Load() HIFCR_Bits   { return HIFCR_Bits(rm.UM32.Load()) }
func (rm HIFCR_Mask) Store(b HIFCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) CTCIF7() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CTCIF7)}}
}

func (p *DMA_Periph) CHTIF7() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CHTIF7)}}
}

func (p *DMA_Periph) CTEIF7() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CTEIF7)}}
}

func (p *DMA_Periph) CDMEIF7() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CDMEIF7)}}
}

func (p *DMA_Periph) CFEIF7() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CFEIF7)}}
}

func (p *DMA_Periph) CTCIF6() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CTCIF6)}}
}

func (p *DMA_Periph) CHTIF6() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CHTIF6)}}
}

func (p *DMA_Periph) CTEIF6() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CTEIF6)}}
}

func (p *DMA_Periph) CDMEIF6() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CDMEIF6)}}
}

func (p *DMA_Periph) CFEIF6() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CFEIF6)}}
}

func (p *DMA_Periph) CTCIF5() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CTCIF5)}}
}

func (p *DMA_Periph) CHTIF5() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CHTIF5)}}
}

func (p *DMA_Periph) CTEIF5() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CTEIF5)}}
}

func (p *DMA_Periph) CDMEIF5() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CDMEIF5)}}
}

func (p *DMA_Periph) CFEIF5() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CFEIF5)}}
}

func (p *DMA_Periph) CTCIF4() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CTCIF4)}}
}

func (p *DMA_Periph) CHTIF4() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CHTIF4)}}
}

func (p *DMA_Periph) CTEIF4() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CTEIF4)}}
}

func (p *DMA_Periph) CDMEIF4() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CDMEIF4)}}
}

func (p *DMA_Periph) CFEIF4() HIFCR_Mask {
	return HIFCR_Mask{mmio.UM32{&p.HIFCR.U32, uint32(CFEIF4)}}
}
