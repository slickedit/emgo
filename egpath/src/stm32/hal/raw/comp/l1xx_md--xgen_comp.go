// +build l1xx_md

package comp

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type COMP_Periph struct {
	CSR CSR
}

func (p *COMP_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var COMP = (*COMP_Periph)(unsafe.Pointer(uintptr(mmap.COMP_BASE)))

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

func (p *COMP_Periph) V10KPU() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(V10KPU)}}
}

func (p *COMP_Periph) V400KPU() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(V400KPU)}}
}

func (p *COMP_Periph) V10KPD() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(V10KPD)}}
}

func (p *COMP_Periph) V400KPD() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(V400KPD)}}
}

func (p *COMP_Periph) CMP1EN() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(CMP1EN)}}
}

func (p *COMP_Periph) SW1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(SW1)}}
}

func (p *COMP_Periph) CMP1OUT() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(CMP1OUT)}}
}

func (p *COMP_Periph) SPEED() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(SPEED)}}
}

func (p *COMP_Periph) CMP2OUT() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(CMP2OUT)}}
}

func (p *COMP_Periph) VREFOUTEN() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(VREFOUTEN)}}
}

func (p *COMP_Periph) WNDWE() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(WNDWE)}}
}

func (p *COMP_Periph) INSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(INSEL)}}
}

func (p *COMP_Periph) OUTSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OUTSEL)}}
}

func (p *COMP_Periph) FCH3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(FCH3)}}
}

func (p *COMP_Periph) FCH8() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(FCH8)}}
}

func (p *COMP_Periph) RCH13() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(RCH13)}}
}

func (p *COMP_Periph) CAIE() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(CAIE)}}
}

func (p *COMP_Periph) CAIF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(CAIF)}}
}

func (p *COMP_Periph) TSUSP() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(TSUSP)}}
}
