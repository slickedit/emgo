package hash

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type HASH_Periph struct {
	CR  CR
	DIN DIN
	STR STR
	HR  [5]HR
	IMR IMR
	SR  SR
	_   [52]uint32
	CSR [54]CSR
}

var HASH = (*HASH_Periph)(unsafe.Pointer(uintptr(mmap.HASH_BASE)))

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

func (p *HASH_Periph) INIT() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(INIT)}}
}

func (p *HASH_Periph) DMAE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(DMAE)}}
}

func (p *HASH_Periph) DATATYPE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(DATATYPE)}}
}

func (p *HASH_Periph) MODE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODE)}}
}

func (p *HASH_Periph) ALGO() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(ALGO)}}
}

func (p *HASH_Periph) NBW() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(NBW)}}
}

func (p *HASH_Periph) DINNE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(DINNE)}}
}

func (p *HASH_Periph) MDMAT() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MDMAT)}}
}

func (p *HASH_Periph) LKEY() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(LKEY)}}
}

type DIN_Bits uint32

func (b DIN_Bits) Field(mask DIN_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIN_Bits) J(v int) DIN_Bits {
	return DIN_Bits(bits.Make32(v, uint32(mask)))
}

type DIN struct{ mmio.U32 }

func (r *DIN) Bits(mask DIN_Bits) DIN_Bits { return DIN_Bits(r.U32.Bits(uint32(mask))) }
func (r *DIN) StoreBits(mask, b DIN_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIN) SetBits(mask DIN_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *DIN) ClearBits(mask DIN_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *DIN) Load() DIN_Bits              { return DIN_Bits(r.U32.Load()) }
func (r *DIN) Store(b DIN_Bits)            { r.U32.Store(uint32(b)) }

type DIN_Mask struct{ mmio.UM32 }

func (rm DIN_Mask) Load() DIN_Bits   { return DIN_Bits(rm.UM32.Load()) }
func (rm DIN_Mask) Store(b DIN_Bits) { rm.UM32.Store(uint32(b)) }

type STR_Bits uint32

func (b STR_Bits) Field(mask STR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask STR_Bits) J(v int) STR_Bits {
	return STR_Bits(bits.Make32(v, uint32(mask)))
}

type STR struct{ mmio.U32 }

func (r *STR) Bits(mask STR_Bits) STR_Bits { return STR_Bits(r.U32.Bits(uint32(mask))) }
func (r *STR) StoreBits(mask, b STR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *STR) SetBits(mask STR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *STR) ClearBits(mask STR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *STR) Load() STR_Bits              { return STR_Bits(r.U32.Load()) }
func (r *STR) Store(b STR_Bits)            { r.U32.Store(uint32(b)) }

type STR_Mask struct{ mmio.UM32 }

func (rm STR_Mask) Load() STR_Bits   { return STR_Bits(rm.UM32.Load()) }
func (rm STR_Mask) Store(b STR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *HASH_Periph) NBW() STR_Mask {
	return STR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(NBW)}}
}

func (p *HASH_Periph) DCAL() STR_Mask {
	return STR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(DCAL)}}
}

type HR_Bits uint32

func (b HR_Bits) Field(mask HR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HR_Bits) J(v int) HR_Bits {
	return HR_Bits(bits.Make32(v, uint32(mask)))
}

type HR struct{ mmio.U32 }

func (r *HR) Bits(mask HR_Bits) HR_Bits { return HR_Bits(r.U32.Bits(uint32(mask))) }
func (r *HR) StoreBits(mask, b HR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HR) SetBits(mask HR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *HR) ClearBits(mask HR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *HR) Load() HR_Bits             { return HR_Bits(r.U32.Load()) }
func (r *HR) Store(b HR_Bits)           { r.U32.Store(uint32(b)) }

type HR_Mask struct{ mmio.UM32 }

func (rm HR_Mask) Load() HR_Bits   { return HR_Bits(rm.UM32.Load()) }
func (rm HR_Mask) Store(b HR_Bits) { rm.UM32.Store(uint32(b)) }

type IMR_Bits uint32

func (b IMR_Bits) Field(mask IMR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IMR_Bits) J(v int) IMR_Bits {
	return IMR_Bits(bits.Make32(v, uint32(mask)))
}

type IMR struct{ mmio.U32 }

func (r *IMR) Bits(mask IMR_Bits) IMR_Bits { return IMR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IMR) StoreBits(mask, b IMR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IMR) SetBits(mask IMR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IMR) ClearBits(mask IMR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IMR) Load() IMR_Bits              { return IMR_Bits(r.U32.Load()) }
func (r *IMR) Store(b IMR_Bits)            { r.U32.Store(uint32(b)) }

type IMR_Mask struct{ mmio.UM32 }

func (rm IMR_Mask) Load() IMR_Bits   { return IMR_Bits(rm.UM32.Load()) }
func (rm IMR_Mask) Store(b IMR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *HASH_Periph) DINIM() IMR_Mask {
	return IMR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 32)), uint32(DINIM)}}
}

func (p *HASH_Periph) DCIM() IMR_Mask {
	return IMR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 32)), uint32(DCIM)}}
}

type SR_Bits uint32

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *HASH_Periph) DINIS() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(DINIS)}}
}

func (p *HASH_Periph) DCIS() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(DCIS)}}
}

func (p *HASH_Periph) DMAS() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(DMAS)}}
}

func (p *HASH_Periph) BUSY() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(BUSY)}}
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
