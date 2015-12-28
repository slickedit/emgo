package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type FSMC_Bank4_Periph struct {
	PCR4  PCR4
	SR4   SR4
	PMEM4 PMEM4
	PATT4 PATT4
	PIO4  PIO4
}

var FSMC_Bank4 = (*FSMC_Bank4_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank4_R_BASE)))

type PCR4_Bits uint32

func (b PCR4_Bits) Field(mask PCR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCR4_Bits) J(v int) PCR4_Bits {
	return PCR4_Bits(bits.Make32(v, uint32(mask)))
}

type PCR4 struct{ mmio.U32 }

func (r *PCR4) Bits(mask PCR4_Bits) PCR4_Bits { return PCR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *PCR4) StoreBits(mask, b PCR4_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PCR4) SetBits(mask PCR4_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PCR4) ClearBits(mask PCR4_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PCR4) Load() PCR4_Bits               { return PCR4_Bits(r.U32.Load()) }
func (r *PCR4) Store(b PCR4_Bits)             { r.U32.Store(uint32(b)) }

type PCR4_Mask struct{ mmio.UM32 }

func (rm PCR4_Mask) Load() PCR4_Bits   { return PCR4_Bits(rm.UM32.Load()) }
func (rm PCR4_Mask) Store(b PCR4_Bits) { rm.UM32.Store(uint32(b)) }

type SR4_Bits uint32

func (b SR4_Bits) Field(mask SR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR4_Bits) J(v int) SR4_Bits {
	return SR4_Bits(bits.Make32(v, uint32(mask)))
}

type SR4 struct{ mmio.U32 }

func (r *SR4) Bits(mask SR4_Bits) SR4_Bits { return SR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR4) StoreBits(mask, b SR4_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR4) SetBits(mask SR4_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *SR4) ClearBits(mask SR4_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *SR4) Load() SR4_Bits              { return SR4_Bits(r.U32.Load()) }
func (r *SR4) Store(b SR4_Bits)            { r.U32.Store(uint32(b)) }

type SR4_Mask struct{ mmio.UM32 }

func (rm SR4_Mask) Load() SR4_Bits   { return SR4_Bits(rm.UM32.Load()) }
func (rm SR4_Mask) Store(b SR4_Bits) { rm.UM32.Store(uint32(b)) }

type PMEM4_Bits uint32

func (b PMEM4_Bits) Field(mask PMEM4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PMEM4_Bits) J(v int) PMEM4_Bits {
	return PMEM4_Bits(bits.Make32(v, uint32(mask)))
}

type PMEM4 struct{ mmio.U32 }

func (r *PMEM4) Bits(mask PMEM4_Bits) PMEM4_Bits { return PMEM4_Bits(r.U32.Bits(uint32(mask))) }
func (r *PMEM4) StoreBits(mask, b PMEM4_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PMEM4) SetBits(mask PMEM4_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PMEM4) ClearBits(mask PMEM4_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PMEM4) Load() PMEM4_Bits                { return PMEM4_Bits(r.U32.Load()) }
func (r *PMEM4) Store(b PMEM4_Bits)              { r.U32.Store(uint32(b)) }

type PMEM4_Mask struct{ mmio.UM32 }

func (rm PMEM4_Mask) Load() PMEM4_Bits   { return PMEM4_Bits(rm.UM32.Load()) }
func (rm PMEM4_Mask) Store(b PMEM4_Bits) { rm.UM32.Store(uint32(b)) }

type PATT4_Bits uint32

func (b PATT4_Bits) Field(mask PATT4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PATT4_Bits) J(v int) PATT4_Bits {
	return PATT4_Bits(bits.Make32(v, uint32(mask)))
}

type PATT4 struct{ mmio.U32 }

func (r *PATT4) Bits(mask PATT4_Bits) PATT4_Bits { return PATT4_Bits(r.U32.Bits(uint32(mask))) }
func (r *PATT4) StoreBits(mask, b PATT4_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PATT4) SetBits(mask PATT4_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PATT4) ClearBits(mask PATT4_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PATT4) Load() PATT4_Bits                { return PATT4_Bits(r.U32.Load()) }
func (r *PATT4) Store(b PATT4_Bits)              { r.U32.Store(uint32(b)) }

type PATT4_Mask struct{ mmio.UM32 }

func (rm PATT4_Mask) Load() PATT4_Bits   { return PATT4_Bits(rm.UM32.Load()) }
func (rm PATT4_Mask) Store(b PATT4_Bits) { rm.UM32.Store(uint32(b)) }

type PIO4_Bits uint32

func (b PIO4_Bits) Field(mask PIO4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PIO4_Bits) J(v int) PIO4_Bits {
	return PIO4_Bits(bits.Make32(v, uint32(mask)))
}

type PIO4 struct{ mmio.U32 }

func (r *PIO4) Bits(mask PIO4_Bits) PIO4_Bits { return PIO4_Bits(r.U32.Bits(uint32(mask))) }
func (r *PIO4) StoreBits(mask, b PIO4_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PIO4) SetBits(mask PIO4_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PIO4) ClearBits(mask PIO4_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PIO4) Load() PIO4_Bits               { return PIO4_Bits(r.U32.Load()) }
func (r *PIO4) Store(b PIO4_Bits)             { r.U32.Store(uint32(b)) }

type PIO4_Mask struct{ mmio.UM32 }

func (rm PIO4_Mask) Load() PIO4_Bits   { return PIO4_Bits(rm.UM32.Load()) }
func (rm PIO4_Mask) Store(b PIO4_Bits) { rm.UM32.Store(uint32(b)) }
