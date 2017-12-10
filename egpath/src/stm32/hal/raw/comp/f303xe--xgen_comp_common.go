// +build f303xe

package comp

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type COMP_Common_Periph struct {
	CSR CSR
}

func (p *COMP_Common_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var COMP12_COMMON = (*COMP_Common_Periph)(unsafe.Pointer(uintptr(mmap.COMP2_BASE)))

//emgo:const
var COMP34_COMMON = (*COMP_Common_Periph)(unsafe.Pointer(uintptr(mmap.COMP4_BASE)))

//emgo:const
var COMP56_COMMON = (*COMP_Common_Periph)(unsafe.Pointer(uintptr(mmap.COMP6_BASE)))

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

func (r *CSR) AtomicStoreBits(mask, b CSR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CSR) AtomicSetBits(mask CSR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CSR) AtomicClearBits(mask CSR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CSR_Mask struct{ mmio.UM32 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM32.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM32.Store(uint32(b)) }