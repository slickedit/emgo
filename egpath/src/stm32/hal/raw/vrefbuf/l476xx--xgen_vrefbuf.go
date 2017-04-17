// +build l476xx

package vrefbuf

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type VREFBUF_Periph struct {
	CSR CSR
	CCR CCR
}

func (p *VREFBUF_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var VREFBUF = (*VREFBUF_Periph)(unsafe.Pointer(uintptr(mmap.VREFBUF_BASE)))

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

func (p *VREFBUF_Periph) ENVR() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(ENVR)}}
}

func (p *VREFBUF_Periph) HIZ() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(HIZ)}}
}

func (p *VREFBUF_Periph) VRS() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(VRS)}}
}

func (p *VREFBUF_Periph) VRR() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(VRR)}}
}

type CCR_Bits uint32

func (b CCR_Bits) Field(mask CCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR_Bits) J(v int) CCR_Bits {
	return CCR_Bits(bits.Make32(v, uint32(mask)))
}

type CCR struct{ mmio.U32 }

func (r *CCR) Bits(mask CCR_Bits) CCR_Bits { return CCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CCR) StoreBits(mask, b CCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CCR) SetBits(mask CCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CCR) ClearBits(mask CCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CCR) Load() CCR_Bits              { return CCR_Bits(r.U32.Load()) }
func (r *CCR) Store(b CCR_Bits)            { r.U32.Store(uint32(b)) }

type CCR_Mask struct{ mmio.UM32 }

func (rm CCR_Mask) Load() CCR_Bits   { return CCR_Bits(rm.UM32.Load()) }
func (rm CCR_Mask) Store(b CCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *VREFBUF_Periph) TRIM() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(TRIM)}}
}
