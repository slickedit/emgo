// +build l476xx

package firewall

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type FIREWALL_Periph struct {
	CSSA   CSSA
	CSL    CSL
	NVDSSA NVDSSA
	NVDSL  NVDSL
	VDSSA  VDSSA
	VDSL   VDSL
	_      [2]uint32
	CR     CR
}

func (p *FIREWALL_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FIREWALL = (*FIREWALL_Periph)(unsafe.Pointer(uintptr(mmap.FIREWALL_BASE)))

type CSSA_Bits uint32

func (b CSSA_Bits) Field(mask CSSA_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSSA_Bits) J(v int) CSSA_Bits {
	return CSSA_Bits(bits.Make32(v, uint32(mask)))
}

type CSSA struct{ mmio.U32 }

func (r *CSSA) Bits(mask CSSA_Bits) CSSA_Bits { return CSSA_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSSA) StoreBits(mask, b CSSA_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSSA) SetBits(mask CSSA_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CSSA) ClearBits(mask CSSA_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CSSA) Load() CSSA_Bits               { return CSSA_Bits(r.U32.Load()) }
func (r *CSSA) Store(b CSSA_Bits)             { r.U32.Store(uint32(b)) }

type CSSA_Mask struct{ mmio.UM32 }

func (rm CSSA_Mask) Load() CSSA_Bits   { return CSSA_Bits(rm.UM32.Load()) }
func (rm CSSA_Mask) Store(b CSSA_Bits) { rm.UM32.Store(uint32(b)) }

type CSL_Bits uint32

func (b CSL_Bits) Field(mask CSL_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSL_Bits) J(v int) CSL_Bits {
	return CSL_Bits(bits.Make32(v, uint32(mask)))
}

type CSL struct{ mmio.U32 }

func (r *CSL) Bits(mask CSL_Bits) CSL_Bits { return CSL_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSL) StoreBits(mask, b CSL_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSL) SetBits(mask CSL_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CSL) ClearBits(mask CSL_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CSL) Load() CSL_Bits              { return CSL_Bits(r.U32.Load()) }
func (r *CSL) Store(b CSL_Bits)            { r.U32.Store(uint32(b)) }

type CSL_Mask struct{ mmio.UM32 }

func (rm CSL_Mask) Load() CSL_Bits   { return CSL_Bits(rm.UM32.Load()) }
func (rm CSL_Mask) Store(b CSL_Bits) { rm.UM32.Store(uint32(b)) }

type NVDSSA_Bits uint32

func (b NVDSSA_Bits) Field(mask NVDSSA_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask NVDSSA_Bits) J(v int) NVDSSA_Bits {
	return NVDSSA_Bits(bits.Make32(v, uint32(mask)))
}

type NVDSSA struct{ mmio.U32 }

func (r *NVDSSA) Bits(mask NVDSSA_Bits) NVDSSA_Bits { return NVDSSA_Bits(r.U32.Bits(uint32(mask))) }
func (r *NVDSSA) StoreBits(mask, b NVDSSA_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *NVDSSA) SetBits(mask NVDSSA_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *NVDSSA) ClearBits(mask NVDSSA_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *NVDSSA) Load() NVDSSA_Bits                 { return NVDSSA_Bits(r.U32.Load()) }
func (r *NVDSSA) Store(b NVDSSA_Bits)               { r.U32.Store(uint32(b)) }

type NVDSSA_Mask struct{ mmio.UM32 }

func (rm NVDSSA_Mask) Load() NVDSSA_Bits   { return NVDSSA_Bits(rm.UM32.Load()) }
func (rm NVDSSA_Mask) Store(b NVDSSA_Bits) { rm.UM32.Store(uint32(b)) }

type NVDSL_Bits uint32

func (b NVDSL_Bits) Field(mask NVDSL_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask NVDSL_Bits) J(v int) NVDSL_Bits {
	return NVDSL_Bits(bits.Make32(v, uint32(mask)))
}

type NVDSL struct{ mmio.U32 }

func (r *NVDSL) Bits(mask NVDSL_Bits) NVDSL_Bits { return NVDSL_Bits(r.U32.Bits(uint32(mask))) }
func (r *NVDSL) StoreBits(mask, b NVDSL_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *NVDSL) SetBits(mask NVDSL_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *NVDSL) ClearBits(mask NVDSL_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *NVDSL) Load() NVDSL_Bits                { return NVDSL_Bits(r.U32.Load()) }
func (r *NVDSL) Store(b NVDSL_Bits)              { r.U32.Store(uint32(b)) }

type NVDSL_Mask struct{ mmio.UM32 }

func (rm NVDSL_Mask) Load() NVDSL_Bits   { return NVDSL_Bits(rm.UM32.Load()) }
func (rm NVDSL_Mask) Store(b NVDSL_Bits) { rm.UM32.Store(uint32(b)) }

type VDSSA_Bits uint32

func (b VDSSA_Bits) Field(mask VDSSA_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask VDSSA_Bits) J(v int) VDSSA_Bits {
	return VDSSA_Bits(bits.Make32(v, uint32(mask)))
}

type VDSSA struct{ mmio.U32 }

func (r *VDSSA) Bits(mask VDSSA_Bits) VDSSA_Bits { return VDSSA_Bits(r.U32.Bits(uint32(mask))) }
func (r *VDSSA) StoreBits(mask, b VDSSA_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *VDSSA) SetBits(mask VDSSA_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *VDSSA) ClearBits(mask VDSSA_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *VDSSA) Load() VDSSA_Bits                { return VDSSA_Bits(r.U32.Load()) }
func (r *VDSSA) Store(b VDSSA_Bits)              { r.U32.Store(uint32(b)) }

type VDSSA_Mask struct{ mmio.UM32 }

func (rm VDSSA_Mask) Load() VDSSA_Bits   { return VDSSA_Bits(rm.UM32.Load()) }
func (rm VDSSA_Mask) Store(b VDSSA_Bits) { rm.UM32.Store(uint32(b)) }

type VDSL_Bits uint32

func (b VDSL_Bits) Field(mask VDSL_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask VDSL_Bits) J(v int) VDSL_Bits {
	return VDSL_Bits(bits.Make32(v, uint32(mask)))
}

type VDSL struct{ mmio.U32 }

func (r *VDSL) Bits(mask VDSL_Bits) VDSL_Bits { return VDSL_Bits(r.U32.Bits(uint32(mask))) }
func (r *VDSL) StoreBits(mask, b VDSL_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *VDSL) SetBits(mask VDSL_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *VDSL) ClearBits(mask VDSL_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *VDSL) Load() VDSL_Bits               { return VDSL_Bits(r.U32.Load()) }
func (r *VDSL) Store(b VDSL_Bits)             { r.U32.Store(uint32(b)) }

type VDSL_Mask struct{ mmio.UM32 }

func (rm VDSL_Mask) Load() VDSL_Bits   { return VDSL_Bits(rm.UM32.Load()) }
func (rm VDSL_Mask) Store(b VDSL_Bits) { rm.UM32.Store(uint32(b)) }

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
