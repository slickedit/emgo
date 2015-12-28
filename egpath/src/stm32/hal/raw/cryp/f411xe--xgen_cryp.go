// +build f411xe

package cryp

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type CRYP_Periph struct {
	CR         CR
	SR         SR
	DR         DR
	DOUT       DOUT
	DMACR      DMACR
	IMSCR      IMSCR
	RISR       RISR
	MISR       MISR
	K0LR       K0LR
	K0RR       K0RR
	K1LR       K1LR
	K1RR       K1RR
	K2LR       K2LR
	K2RR       K2RR
	K3LR       K3LR
	K3RR       K3RR
	IV0LR      IV0LR
	IV0RR      IV0RR
	IV1LR      IV1LR
	IV1RR      IV1RR
	CSGCMCCM0R CSGCMCCM0R
	CSGCMCCM1R CSGCMCCM1R
	CSGCMCCM2R CSGCMCCM2R
	CSGCMCCM3R CSGCMCCM3R
	CSGCMCCM4R CSGCMCCM4R
	CSGCMCCM5R CSGCMCCM5R
	CSGCMCCM6R CSGCMCCM6R
	CSGCMCCM7R CSGCMCCM7R
	CSGCM0R    CSGCM0R
	CSGCM1R    CSGCM1R
	CSGCM2R    CSGCM2R
	CSGCM3R    CSGCM3R
	CSGCM4R    CSGCM4R
	CSGCM5R    CSGCM5R
	CSGCM6R    CSGCM6R
	CSGCM7R    CSGCM7R
}

var CRYP = (*CRYP_Periph)(unsafe.Pointer(uintptr(mmap.CRYP_BASE)))

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

func (p *CRYP_Periph) ALGODIR() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(ALGODIR)}}
}

func (p *CRYP_Periph) ALGOMODE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(ALGOMODE)}}
}

func (p *CRYP_Periph) DATATYPE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(DATATYPE)}}
}

func (p *CRYP_Periph) KEYSIZE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(KEYSIZE)}}
}

func (p *CRYP_Periph) FFLUSH() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(FFLUSH)}}
}

func (p *CRYP_Periph) CRYPEN() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(CRYPEN)}}
}

func (p *CRYP_Periph) GCM_CCMPH() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(GCM_CCMPH)}}
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

func (p *CRYP_Periph) IFEM() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(IFEM)}}
}

func (p *CRYP_Periph) IFNF() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(IFNF)}}
}

func (p *CRYP_Periph) OFNE() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OFNE)}}
}

func (p *CRYP_Periph) OFFU() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(OFFU)}}
}

func (p *CRYP_Periph) BUSY() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(BUSY)}}
}

type DR_Bits uint32

func (b DR_Bits) Field(mask DR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR_Bits) J(v int) DR_Bits {
	return DR_Bits(bits.Make32(v, uint32(mask)))
}

type DR struct{ mmio.U32 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U32.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U32.Store(uint32(b)) }

type DR_Mask struct{ mmio.UM32 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM32.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM32.Store(uint32(b)) }

type DOUT_Bits uint32

func (b DOUT_Bits) Field(mask DOUT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOUT_Bits) J(v int) DOUT_Bits {
	return DOUT_Bits(bits.Make32(v, uint32(mask)))
}

type DOUT struct{ mmio.U32 }

func (r *DOUT) Bits(mask DOUT_Bits) DOUT_Bits { return DOUT_Bits(r.U32.Bits(uint32(mask))) }
func (r *DOUT) StoreBits(mask, b DOUT_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DOUT) SetBits(mask DOUT_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DOUT) ClearBits(mask DOUT_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DOUT) Load() DOUT_Bits               { return DOUT_Bits(r.U32.Load()) }
func (r *DOUT) Store(b DOUT_Bits)             { r.U32.Store(uint32(b)) }

type DOUT_Mask struct{ mmio.UM32 }

func (rm DOUT_Mask) Load() DOUT_Bits   { return DOUT_Bits(rm.UM32.Load()) }
func (rm DOUT_Mask) Store(b DOUT_Bits) { rm.UM32.Store(uint32(b)) }

type DMACR_Bits uint32

func (b DMACR_Bits) Field(mask DMACR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DMACR_Bits) J(v int) DMACR_Bits {
	return DMACR_Bits(bits.Make32(v, uint32(mask)))
}

type DMACR struct{ mmio.U32 }

func (r *DMACR) Bits(mask DMACR_Bits) DMACR_Bits { return DMACR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DMACR) StoreBits(mask, b DMACR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DMACR) SetBits(mask DMACR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *DMACR) ClearBits(mask DMACR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *DMACR) Load() DMACR_Bits                { return DMACR_Bits(r.U32.Load()) }
func (r *DMACR) Store(b DMACR_Bits)              { r.U32.Store(uint32(b)) }

type DMACR_Mask struct{ mmio.UM32 }

func (rm DMACR_Mask) Load() DMACR_Bits   { return DMACR_Bits(rm.UM32.Load()) }
func (rm DMACR_Mask) Store(b DMACR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CRYP_Periph) DIEN() DMACR_Mask {
	return DMACR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(DIEN)}}
}

func (p *CRYP_Periph) DOEN() DMACR_Mask {
	return DMACR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(DOEN)}}
}

type IMSCR_Bits uint32

func (b IMSCR_Bits) Field(mask IMSCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IMSCR_Bits) J(v int) IMSCR_Bits {
	return IMSCR_Bits(bits.Make32(v, uint32(mask)))
}

type IMSCR struct{ mmio.U32 }

func (r *IMSCR) Bits(mask IMSCR_Bits) IMSCR_Bits { return IMSCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IMSCR) StoreBits(mask, b IMSCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IMSCR) SetBits(mask IMSCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *IMSCR) ClearBits(mask IMSCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *IMSCR) Load() IMSCR_Bits                { return IMSCR_Bits(r.U32.Load()) }
func (r *IMSCR) Store(b IMSCR_Bits)              { r.U32.Store(uint32(b)) }

type IMSCR_Mask struct{ mmio.UM32 }

func (rm IMSCR_Mask) Load() IMSCR_Bits   { return IMSCR_Bits(rm.UM32.Load()) }
func (rm IMSCR_Mask) Store(b IMSCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CRYP_Periph) INIM() IMSCR_Mask {
	return IMSCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(INIM)}}
}

func (p *CRYP_Periph) OUTIM() IMSCR_Mask {
	return IMSCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(OUTIM)}}
}

type RISR_Bits uint32

func (b RISR_Bits) Field(mask RISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RISR_Bits) J(v int) RISR_Bits {
	return RISR_Bits(bits.Make32(v, uint32(mask)))
}

type RISR struct{ mmio.U32 }

func (r *RISR) Bits(mask RISR_Bits) RISR_Bits { return RISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RISR) StoreBits(mask, b RISR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RISR) SetBits(mask RISR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *RISR) ClearBits(mask RISR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *RISR) Load() RISR_Bits               { return RISR_Bits(r.U32.Load()) }
func (r *RISR) Store(b RISR_Bits)             { r.U32.Store(uint32(b)) }

type RISR_Mask struct{ mmio.UM32 }

func (rm RISR_Mask) Load() RISR_Bits   { return RISR_Bits(rm.UM32.Load()) }
func (rm RISR_Mask) Store(b RISR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CRYP_Periph) OUTRIS() RISR_Mask {
	return RISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(OUTRIS)}}
}

func (p *CRYP_Periph) INRIS() RISR_Mask {
	return RISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(INRIS)}}
}

type MISR_Bits uint32

func (b MISR_Bits) Field(mask MISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MISR_Bits) J(v int) MISR_Bits {
	return MISR_Bits(bits.Make32(v, uint32(mask)))
}

type MISR struct{ mmio.U32 }

func (r *MISR) Bits(mask MISR_Bits) MISR_Bits { return MISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *MISR) StoreBits(mask, b MISR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MISR) SetBits(mask MISR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *MISR) ClearBits(mask MISR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *MISR) Load() MISR_Bits               { return MISR_Bits(r.U32.Load()) }
func (r *MISR) Store(b MISR_Bits)             { r.U32.Store(uint32(b)) }

type MISR_Mask struct{ mmio.UM32 }

func (rm MISR_Mask) Load() MISR_Bits   { return MISR_Bits(rm.UM32.Load()) }
func (rm MISR_Mask) Store(b MISR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CRYP_Periph) INMIS() MISR_Mask {
	return MISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(INMIS)}}
}

func (p *CRYP_Periph) OUTMIS() MISR_Mask {
	return MISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(OUTMIS)}}
}

type K0LR_Bits uint32

func (b K0LR_Bits) Field(mask K0LR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask K0LR_Bits) J(v int) K0LR_Bits {
	return K0LR_Bits(bits.Make32(v, uint32(mask)))
}

type K0LR struct{ mmio.U32 }

func (r *K0LR) Bits(mask K0LR_Bits) K0LR_Bits { return K0LR_Bits(r.U32.Bits(uint32(mask))) }
func (r *K0LR) StoreBits(mask, b K0LR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *K0LR) SetBits(mask K0LR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *K0LR) ClearBits(mask K0LR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *K0LR) Load() K0LR_Bits               { return K0LR_Bits(r.U32.Load()) }
func (r *K0LR) Store(b K0LR_Bits)             { r.U32.Store(uint32(b)) }

type K0LR_Mask struct{ mmio.UM32 }

func (rm K0LR_Mask) Load() K0LR_Bits   { return K0LR_Bits(rm.UM32.Load()) }
func (rm K0LR_Mask) Store(b K0LR_Bits) { rm.UM32.Store(uint32(b)) }

type K0RR_Bits uint32

func (b K0RR_Bits) Field(mask K0RR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask K0RR_Bits) J(v int) K0RR_Bits {
	return K0RR_Bits(bits.Make32(v, uint32(mask)))
}

type K0RR struct{ mmio.U32 }

func (r *K0RR) Bits(mask K0RR_Bits) K0RR_Bits { return K0RR_Bits(r.U32.Bits(uint32(mask))) }
func (r *K0RR) StoreBits(mask, b K0RR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *K0RR) SetBits(mask K0RR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *K0RR) ClearBits(mask K0RR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *K0RR) Load() K0RR_Bits               { return K0RR_Bits(r.U32.Load()) }
func (r *K0RR) Store(b K0RR_Bits)             { r.U32.Store(uint32(b)) }

type K0RR_Mask struct{ mmio.UM32 }

func (rm K0RR_Mask) Load() K0RR_Bits   { return K0RR_Bits(rm.UM32.Load()) }
func (rm K0RR_Mask) Store(b K0RR_Bits) { rm.UM32.Store(uint32(b)) }

type K1LR_Bits uint32

func (b K1LR_Bits) Field(mask K1LR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask K1LR_Bits) J(v int) K1LR_Bits {
	return K1LR_Bits(bits.Make32(v, uint32(mask)))
}

type K1LR struct{ mmio.U32 }

func (r *K1LR) Bits(mask K1LR_Bits) K1LR_Bits { return K1LR_Bits(r.U32.Bits(uint32(mask))) }
func (r *K1LR) StoreBits(mask, b K1LR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *K1LR) SetBits(mask K1LR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *K1LR) ClearBits(mask K1LR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *K1LR) Load() K1LR_Bits               { return K1LR_Bits(r.U32.Load()) }
func (r *K1LR) Store(b K1LR_Bits)             { r.U32.Store(uint32(b)) }

type K1LR_Mask struct{ mmio.UM32 }

func (rm K1LR_Mask) Load() K1LR_Bits   { return K1LR_Bits(rm.UM32.Load()) }
func (rm K1LR_Mask) Store(b K1LR_Bits) { rm.UM32.Store(uint32(b)) }

type K1RR_Bits uint32

func (b K1RR_Bits) Field(mask K1RR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask K1RR_Bits) J(v int) K1RR_Bits {
	return K1RR_Bits(bits.Make32(v, uint32(mask)))
}

type K1RR struct{ mmio.U32 }

func (r *K1RR) Bits(mask K1RR_Bits) K1RR_Bits { return K1RR_Bits(r.U32.Bits(uint32(mask))) }
func (r *K1RR) StoreBits(mask, b K1RR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *K1RR) SetBits(mask K1RR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *K1RR) ClearBits(mask K1RR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *K1RR) Load() K1RR_Bits               { return K1RR_Bits(r.U32.Load()) }
func (r *K1RR) Store(b K1RR_Bits)             { r.U32.Store(uint32(b)) }

type K1RR_Mask struct{ mmio.UM32 }

func (rm K1RR_Mask) Load() K1RR_Bits   { return K1RR_Bits(rm.UM32.Load()) }
func (rm K1RR_Mask) Store(b K1RR_Bits) { rm.UM32.Store(uint32(b)) }

type K2LR_Bits uint32

func (b K2LR_Bits) Field(mask K2LR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask K2LR_Bits) J(v int) K2LR_Bits {
	return K2LR_Bits(bits.Make32(v, uint32(mask)))
}

type K2LR struct{ mmio.U32 }

func (r *K2LR) Bits(mask K2LR_Bits) K2LR_Bits { return K2LR_Bits(r.U32.Bits(uint32(mask))) }
func (r *K2LR) StoreBits(mask, b K2LR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *K2LR) SetBits(mask K2LR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *K2LR) ClearBits(mask K2LR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *K2LR) Load() K2LR_Bits               { return K2LR_Bits(r.U32.Load()) }
func (r *K2LR) Store(b K2LR_Bits)             { r.U32.Store(uint32(b)) }

type K2LR_Mask struct{ mmio.UM32 }

func (rm K2LR_Mask) Load() K2LR_Bits   { return K2LR_Bits(rm.UM32.Load()) }
func (rm K2LR_Mask) Store(b K2LR_Bits) { rm.UM32.Store(uint32(b)) }

type K2RR_Bits uint32

func (b K2RR_Bits) Field(mask K2RR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask K2RR_Bits) J(v int) K2RR_Bits {
	return K2RR_Bits(bits.Make32(v, uint32(mask)))
}

type K2RR struct{ mmio.U32 }

func (r *K2RR) Bits(mask K2RR_Bits) K2RR_Bits { return K2RR_Bits(r.U32.Bits(uint32(mask))) }
func (r *K2RR) StoreBits(mask, b K2RR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *K2RR) SetBits(mask K2RR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *K2RR) ClearBits(mask K2RR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *K2RR) Load() K2RR_Bits               { return K2RR_Bits(r.U32.Load()) }
func (r *K2RR) Store(b K2RR_Bits)             { r.U32.Store(uint32(b)) }

type K2RR_Mask struct{ mmio.UM32 }

func (rm K2RR_Mask) Load() K2RR_Bits   { return K2RR_Bits(rm.UM32.Load()) }
func (rm K2RR_Mask) Store(b K2RR_Bits) { rm.UM32.Store(uint32(b)) }

type K3LR_Bits uint32

func (b K3LR_Bits) Field(mask K3LR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask K3LR_Bits) J(v int) K3LR_Bits {
	return K3LR_Bits(bits.Make32(v, uint32(mask)))
}

type K3LR struct{ mmio.U32 }

func (r *K3LR) Bits(mask K3LR_Bits) K3LR_Bits { return K3LR_Bits(r.U32.Bits(uint32(mask))) }
func (r *K3LR) StoreBits(mask, b K3LR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *K3LR) SetBits(mask K3LR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *K3LR) ClearBits(mask K3LR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *K3LR) Load() K3LR_Bits               { return K3LR_Bits(r.U32.Load()) }
func (r *K3LR) Store(b K3LR_Bits)             { r.U32.Store(uint32(b)) }

type K3LR_Mask struct{ mmio.UM32 }

func (rm K3LR_Mask) Load() K3LR_Bits   { return K3LR_Bits(rm.UM32.Load()) }
func (rm K3LR_Mask) Store(b K3LR_Bits) { rm.UM32.Store(uint32(b)) }

type K3RR_Bits uint32

func (b K3RR_Bits) Field(mask K3RR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask K3RR_Bits) J(v int) K3RR_Bits {
	return K3RR_Bits(bits.Make32(v, uint32(mask)))
}

type K3RR struct{ mmio.U32 }

func (r *K3RR) Bits(mask K3RR_Bits) K3RR_Bits { return K3RR_Bits(r.U32.Bits(uint32(mask))) }
func (r *K3RR) StoreBits(mask, b K3RR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *K3RR) SetBits(mask K3RR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *K3RR) ClearBits(mask K3RR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *K3RR) Load() K3RR_Bits               { return K3RR_Bits(r.U32.Load()) }
func (r *K3RR) Store(b K3RR_Bits)             { r.U32.Store(uint32(b)) }

type K3RR_Mask struct{ mmio.UM32 }

func (rm K3RR_Mask) Load() K3RR_Bits   { return K3RR_Bits(rm.UM32.Load()) }
func (rm K3RR_Mask) Store(b K3RR_Bits) { rm.UM32.Store(uint32(b)) }

type IV0LR_Bits uint32

func (b IV0LR_Bits) Field(mask IV0LR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IV0LR_Bits) J(v int) IV0LR_Bits {
	return IV0LR_Bits(bits.Make32(v, uint32(mask)))
}

type IV0LR struct{ mmio.U32 }

func (r *IV0LR) Bits(mask IV0LR_Bits) IV0LR_Bits { return IV0LR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IV0LR) StoreBits(mask, b IV0LR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IV0LR) SetBits(mask IV0LR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *IV0LR) ClearBits(mask IV0LR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *IV0LR) Load() IV0LR_Bits                { return IV0LR_Bits(r.U32.Load()) }
func (r *IV0LR) Store(b IV0LR_Bits)              { r.U32.Store(uint32(b)) }

type IV0LR_Mask struct{ mmio.UM32 }

func (rm IV0LR_Mask) Load() IV0LR_Bits   { return IV0LR_Bits(rm.UM32.Load()) }
func (rm IV0LR_Mask) Store(b IV0LR_Bits) { rm.UM32.Store(uint32(b)) }

type IV0RR_Bits uint32

func (b IV0RR_Bits) Field(mask IV0RR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IV0RR_Bits) J(v int) IV0RR_Bits {
	return IV0RR_Bits(bits.Make32(v, uint32(mask)))
}

type IV0RR struct{ mmio.U32 }

func (r *IV0RR) Bits(mask IV0RR_Bits) IV0RR_Bits { return IV0RR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IV0RR) StoreBits(mask, b IV0RR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IV0RR) SetBits(mask IV0RR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *IV0RR) ClearBits(mask IV0RR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *IV0RR) Load() IV0RR_Bits                { return IV0RR_Bits(r.U32.Load()) }
func (r *IV0RR) Store(b IV0RR_Bits)              { r.U32.Store(uint32(b)) }

type IV0RR_Mask struct{ mmio.UM32 }

func (rm IV0RR_Mask) Load() IV0RR_Bits   { return IV0RR_Bits(rm.UM32.Load()) }
func (rm IV0RR_Mask) Store(b IV0RR_Bits) { rm.UM32.Store(uint32(b)) }

type IV1LR_Bits uint32

func (b IV1LR_Bits) Field(mask IV1LR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IV1LR_Bits) J(v int) IV1LR_Bits {
	return IV1LR_Bits(bits.Make32(v, uint32(mask)))
}

type IV1LR struct{ mmio.U32 }

func (r *IV1LR) Bits(mask IV1LR_Bits) IV1LR_Bits { return IV1LR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IV1LR) StoreBits(mask, b IV1LR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IV1LR) SetBits(mask IV1LR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *IV1LR) ClearBits(mask IV1LR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *IV1LR) Load() IV1LR_Bits                { return IV1LR_Bits(r.U32.Load()) }
func (r *IV1LR) Store(b IV1LR_Bits)              { r.U32.Store(uint32(b)) }

type IV1LR_Mask struct{ mmio.UM32 }

func (rm IV1LR_Mask) Load() IV1LR_Bits   { return IV1LR_Bits(rm.UM32.Load()) }
func (rm IV1LR_Mask) Store(b IV1LR_Bits) { rm.UM32.Store(uint32(b)) }

type IV1RR_Bits uint32

func (b IV1RR_Bits) Field(mask IV1RR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IV1RR_Bits) J(v int) IV1RR_Bits {
	return IV1RR_Bits(bits.Make32(v, uint32(mask)))
}

type IV1RR struct{ mmio.U32 }

func (r *IV1RR) Bits(mask IV1RR_Bits) IV1RR_Bits { return IV1RR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IV1RR) StoreBits(mask, b IV1RR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IV1RR) SetBits(mask IV1RR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *IV1RR) ClearBits(mask IV1RR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *IV1RR) Load() IV1RR_Bits                { return IV1RR_Bits(r.U32.Load()) }
func (r *IV1RR) Store(b IV1RR_Bits)              { r.U32.Store(uint32(b)) }

type IV1RR_Mask struct{ mmio.UM32 }

func (rm IV1RR_Mask) Load() IV1RR_Bits   { return IV1RR_Bits(rm.UM32.Load()) }
func (rm IV1RR_Mask) Store(b IV1RR_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCMCCM0R_Bits uint32

func (b CSGCMCCM0R_Bits) Field(mask CSGCMCCM0R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCMCCM0R_Bits) J(v int) CSGCMCCM0R_Bits {
	return CSGCMCCM0R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCMCCM0R struct{ mmio.U32 }

func (r *CSGCMCCM0R) Bits(mask CSGCMCCM0R_Bits) CSGCMCCM0R_Bits {
	return CSGCMCCM0R_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CSGCMCCM0R) StoreBits(mask, b CSGCMCCM0R_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCMCCM0R) SetBits(mask CSGCMCCM0R_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CSGCMCCM0R) ClearBits(mask CSGCMCCM0R_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCMCCM0R) Load() CSGCMCCM0R_Bits             { return CSGCMCCM0R_Bits(r.U32.Load()) }
func (r *CSGCMCCM0R) Store(b CSGCMCCM0R_Bits)           { r.U32.Store(uint32(b)) }

type CSGCMCCM0R_Mask struct{ mmio.UM32 }

func (rm CSGCMCCM0R_Mask) Load() CSGCMCCM0R_Bits   { return CSGCMCCM0R_Bits(rm.UM32.Load()) }
func (rm CSGCMCCM0R_Mask) Store(b CSGCMCCM0R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCMCCM1R_Bits uint32

func (b CSGCMCCM1R_Bits) Field(mask CSGCMCCM1R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCMCCM1R_Bits) J(v int) CSGCMCCM1R_Bits {
	return CSGCMCCM1R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCMCCM1R struct{ mmio.U32 }

func (r *CSGCMCCM1R) Bits(mask CSGCMCCM1R_Bits) CSGCMCCM1R_Bits {
	return CSGCMCCM1R_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CSGCMCCM1R) StoreBits(mask, b CSGCMCCM1R_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCMCCM1R) SetBits(mask CSGCMCCM1R_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CSGCMCCM1R) ClearBits(mask CSGCMCCM1R_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCMCCM1R) Load() CSGCMCCM1R_Bits             { return CSGCMCCM1R_Bits(r.U32.Load()) }
func (r *CSGCMCCM1R) Store(b CSGCMCCM1R_Bits)           { r.U32.Store(uint32(b)) }

type CSGCMCCM1R_Mask struct{ mmio.UM32 }

func (rm CSGCMCCM1R_Mask) Load() CSGCMCCM1R_Bits   { return CSGCMCCM1R_Bits(rm.UM32.Load()) }
func (rm CSGCMCCM1R_Mask) Store(b CSGCMCCM1R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCMCCM2R_Bits uint32

func (b CSGCMCCM2R_Bits) Field(mask CSGCMCCM2R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCMCCM2R_Bits) J(v int) CSGCMCCM2R_Bits {
	return CSGCMCCM2R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCMCCM2R struct{ mmio.U32 }

func (r *CSGCMCCM2R) Bits(mask CSGCMCCM2R_Bits) CSGCMCCM2R_Bits {
	return CSGCMCCM2R_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CSGCMCCM2R) StoreBits(mask, b CSGCMCCM2R_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCMCCM2R) SetBits(mask CSGCMCCM2R_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CSGCMCCM2R) ClearBits(mask CSGCMCCM2R_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCMCCM2R) Load() CSGCMCCM2R_Bits             { return CSGCMCCM2R_Bits(r.U32.Load()) }
func (r *CSGCMCCM2R) Store(b CSGCMCCM2R_Bits)           { r.U32.Store(uint32(b)) }

type CSGCMCCM2R_Mask struct{ mmio.UM32 }

func (rm CSGCMCCM2R_Mask) Load() CSGCMCCM2R_Bits   { return CSGCMCCM2R_Bits(rm.UM32.Load()) }
func (rm CSGCMCCM2R_Mask) Store(b CSGCMCCM2R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCMCCM3R_Bits uint32

func (b CSGCMCCM3R_Bits) Field(mask CSGCMCCM3R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCMCCM3R_Bits) J(v int) CSGCMCCM3R_Bits {
	return CSGCMCCM3R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCMCCM3R struct{ mmio.U32 }

func (r *CSGCMCCM3R) Bits(mask CSGCMCCM3R_Bits) CSGCMCCM3R_Bits {
	return CSGCMCCM3R_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CSGCMCCM3R) StoreBits(mask, b CSGCMCCM3R_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCMCCM3R) SetBits(mask CSGCMCCM3R_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CSGCMCCM3R) ClearBits(mask CSGCMCCM3R_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCMCCM3R) Load() CSGCMCCM3R_Bits             { return CSGCMCCM3R_Bits(r.U32.Load()) }
func (r *CSGCMCCM3R) Store(b CSGCMCCM3R_Bits)           { r.U32.Store(uint32(b)) }

type CSGCMCCM3R_Mask struct{ mmio.UM32 }

func (rm CSGCMCCM3R_Mask) Load() CSGCMCCM3R_Bits   { return CSGCMCCM3R_Bits(rm.UM32.Load()) }
func (rm CSGCMCCM3R_Mask) Store(b CSGCMCCM3R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCMCCM4R_Bits uint32

func (b CSGCMCCM4R_Bits) Field(mask CSGCMCCM4R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCMCCM4R_Bits) J(v int) CSGCMCCM4R_Bits {
	return CSGCMCCM4R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCMCCM4R struct{ mmio.U32 }

func (r *CSGCMCCM4R) Bits(mask CSGCMCCM4R_Bits) CSGCMCCM4R_Bits {
	return CSGCMCCM4R_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CSGCMCCM4R) StoreBits(mask, b CSGCMCCM4R_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCMCCM4R) SetBits(mask CSGCMCCM4R_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CSGCMCCM4R) ClearBits(mask CSGCMCCM4R_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCMCCM4R) Load() CSGCMCCM4R_Bits             { return CSGCMCCM4R_Bits(r.U32.Load()) }
func (r *CSGCMCCM4R) Store(b CSGCMCCM4R_Bits)           { r.U32.Store(uint32(b)) }

type CSGCMCCM4R_Mask struct{ mmio.UM32 }

func (rm CSGCMCCM4R_Mask) Load() CSGCMCCM4R_Bits   { return CSGCMCCM4R_Bits(rm.UM32.Load()) }
func (rm CSGCMCCM4R_Mask) Store(b CSGCMCCM4R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCMCCM5R_Bits uint32

func (b CSGCMCCM5R_Bits) Field(mask CSGCMCCM5R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCMCCM5R_Bits) J(v int) CSGCMCCM5R_Bits {
	return CSGCMCCM5R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCMCCM5R struct{ mmio.U32 }

func (r *CSGCMCCM5R) Bits(mask CSGCMCCM5R_Bits) CSGCMCCM5R_Bits {
	return CSGCMCCM5R_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CSGCMCCM5R) StoreBits(mask, b CSGCMCCM5R_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCMCCM5R) SetBits(mask CSGCMCCM5R_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CSGCMCCM5R) ClearBits(mask CSGCMCCM5R_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCMCCM5R) Load() CSGCMCCM5R_Bits             { return CSGCMCCM5R_Bits(r.U32.Load()) }
func (r *CSGCMCCM5R) Store(b CSGCMCCM5R_Bits)           { r.U32.Store(uint32(b)) }

type CSGCMCCM5R_Mask struct{ mmio.UM32 }

func (rm CSGCMCCM5R_Mask) Load() CSGCMCCM5R_Bits   { return CSGCMCCM5R_Bits(rm.UM32.Load()) }
func (rm CSGCMCCM5R_Mask) Store(b CSGCMCCM5R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCMCCM6R_Bits uint32

func (b CSGCMCCM6R_Bits) Field(mask CSGCMCCM6R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCMCCM6R_Bits) J(v int) CSGCMCCM6R_Bits {
	return CSGCMCCM6R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCMCCM6R struct{ mmio.U32 }

func (r *CSGCMCCM6R) Bits(mask CSGCMCCM6R_Bits) CSGCMCCM6R_Bits {
	return CSGCMCCM6R_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CSGCMCCM6R) StoreBits(mask, b CSGCMCCM6R_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCMCCM6R) SetBits(mask CSGCMCCM6R_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CSGCMCCM6R) ClearBits(mask CSGCMCCM6R_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCMCCM6R) Load() CSGCMCCM6R_Bits             { return CSGCMCCM6R_Bits(r.U32.Load()) }
func (r *CSGCMCCM6R) Store(b CSGCMCCM6R_Bits)           { r.U32.Store(uint32(b)) }

type CSGCMCCM6R_Mask struct{ mmio.UM32 }

func (rm CSGCMCCM6R_Mask) Load() CSGCMCCM6R_Bits   { return CSGCMCCM6R_Bits(rm.UM32.Load()) }
func (rm CSGCMCCM6R_Mask) Store(b CSGCMCCM6R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCMCCM7R_Bits uint32

func (b CSGCMCCM7R_Bits) Field(mask CSGCMCCM7R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCMCCM7R_Bits) J(v int) CSGCMCCM7R_Bits {
	return CSGCMCCM7R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCMCCM7R struct{ mmio.U32 }

func (r *CSGCMCCM7R) Bits(mask CSGCMCCM7R_Bits) CSGCMCCM7R_Bits {
	return CSGCMCCM7R_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CSGCMCCM7R) StoreBits(mask, b CSGCMCCM7R_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCMCCM7R) SetBits(mask CSGCMCCM7R_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CSGCMCCM7R) ClearBits(mask CSGCMCCM7R_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCMCCM7R) Load() CSGCMCCM7R_Bits             { return CSGCMCCM7R_Bits(r.U32.Load()) }
func (r *CSGCMCCM7R) Store(b CSGCMCCM7R_Bits)           { r.U32.Store(uint32(b)) }

type CSGCMCCM7R_Mask struct{ mmio.UM32 }

func (rm CSGCMCCM7R_Mask) Load() CSGCMCCM7R_Bits   { return CSGCMCCM7R_Bits(rm.UM32.Load()) }
func (rm CSGCMCCM7R_Mask) Store(b CSGCMCCM7R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCM0R_Bits uint32

func (b CSGCM0R_Bits) Field(mask CSGCM0R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCM0R_Bits) J(v int) CSGCM0R_Bits {
	return CSGCM0R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCM0R struct{ mmio.U32 }

func (r *CSGCM0R) Bits(mask CSGCM0R_Bits) CSGCM0R_Bits { return CSGCM0R_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSGCM0R) StoreBits(mask, b CSGCM0R_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCM0R) SetBits(mask CSGCM0R_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CSGCM0R) ClearBits(mask CSGCM0R_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCM0R) Load() CSGCM0R_Bits                  { return CSGCM0R_Bits(r.U32.Load()) }
func (r *CSGCM0R) Store(b CSGCM0R_Bits)                { r.U32.Store(uint32(b)) }

type CSGCM0R_Mask struct{ mmio.UM32 }

func (rm CSGCM0R_Mask) Load() CSGCM0R_Bits   { return CSGCM0R_Bits(rm.UM32.Load()) }
func (rm CSGCM0R_Mask) Store(b CSGCM0R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCM1R_Bits uint32

func (b CSGCM1R_Bits) Field(mask CSGCM1R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCM1R_Bits) J(v int) CSGCM1R_Bits {
	return CSGCM1R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCM1R struct{ mmio.U32 }

func (r *CSGCM1R) Bits(mask CSGCM1R_Bits) CSGCM1R_Bits { return CSGCM1R_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSGCM1R) StoreBits(mask, b CSGCM1R_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCM1R) SetBits(mask CSGCM1R_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CSGCM1R) ClearBits(mask CSGCM1R_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCM1R) Load() CSGCM1R_Bits                  { return CSGCM1R_Bits(r.U32.Load()) }
func (r *CSGCM1R) Store(b CSGCM1R_Bits)                { r.U32.Store(uint32(b)) }

type CSGCM1R_Mask struct{ mmio.UM32 }

func (rm CSGCM1R_Mask) Load() CSGCM1R_Bits   { return CSGCM1R_Bits(rm.UM32.Load()) }
func (rm CSGCM1R_Mask) Store(b CSGCM1R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCM2R_Bits uint32

func (b CSGCM2R_Bits) Field(mask CSGCM2R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCM2R_Bits) J(v int) CSGCM2R_Bits {
	return CSGCM2R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCM2R struct{ mmio.U32 }

func (r *CSGCM2R) Bits(mask CSGCM2R_Bits) CSGCM2R_Bits { return CSGCM2R_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSGCM2R) StoreBits(mask, b CSGCM2R_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCM2R) SetBits(mask CSGCM2R_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CSGCM2R) ClearBits(mask CSGCM2R_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCM2R) Load() CSGCM2R_Bits                  { return CSGCM2R_Bits(r.U32.Load()) }
func (r *CSGCM2R) Store(b CSGCM2R_Bits)                { r.U32.Store(uint32(b)) }

type CSGCM2R_Mask struct{ mmio.UM32 }

func (rm CSGCM2R_Mask) Load() CSGCM2R_Bits   { return CSGCM2R_Bits(rm.UM32.Load()) }
func (rm CSGCM2R_Mask) Store(b CSGCM2R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCM3R_Bits uint32

func (b CSGCM3R_Bits) Field(mask CSGCM3R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCM3R_Bits) J(v int) CSGCM3R_Bits {
	return CSGCM3R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCM3R struct{ mmio.U32 }

func (r *CSGCM3R) Bits(mask CSGCM3R_Bits) CSGCM3R_Bits { return CSGCM3R_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSGCM3R) StoreBits(mask, b CSGCM3R_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCM3R) SetBits(mask CSGCM3R_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CSGCM3R) ClearBits(mask CSGCM3R_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCM3R) Load() CSGCM3R_Bits                  { return CSGCM3R_Bits(r.U32.Load()) }
func (r *CSGCM3R) Store(b CSGCM3R_Bits)                { r.U32.Store(uint32(b)) }

type CSGCM3R_Mask struct{ mmio.UM32 }

func (rm CSGCM3R_Mask) Load() CSGCM3R_Bits   { return CSGCM3R_Bits(rm.UM32.Load()) }
func (rm CSGCM3R_Mask) Store(b CSGCM3R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCM4R_Bits uint32

func (b CSGCM4R_Bits) Field(mask CSGCM4R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCM4R_Bits) J(v int) CSGCM4R_Bits {
	return CSGCM4R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCM4R struct{ mmio.U32 }

func (r *CSGCM4R) Bits(mask CSGCM4R_Bits) CSGCM4R_Bits { return CSGCM4R_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSGCM4R) StoreBits(mask, b CSGCM4R_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCM4R) SetBits(mask CSGCM4R_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CSGCM4R) ClearBits(mask CSGCM4R_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCM4R) Load() CSGCM4R_Bits                  { return CSGCM4R_Bits(r.U32.Load()) }
func (r *CSGCM4R) Store(b CSGCM4R_Bits)                { r.U32.Store(uint32(b)) }

type CSGCM4R_Mask struct{ mmio.UM32 }

func (rm CSGCM4R_Mask) Load() CSGCM4R_Bits   { return CSGCM4R_Bits(rm.UM32.Load()) }
func (rm CSGCM4R_Mask) Store(b CSGCM4R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCM5R_Bits uint32

func (b CSGCM5R_Bits) Field(mask CSGCM5R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCM5R_Bits) J(v int) CSGCM5R_Bits {
	return CSGCM5R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCM5R struct{ mmio.U32 }

func (r *CSGCM5R) Bits(mask CSGCM5R_Bits) CSGCM5R_Bits { return CSGCM5R_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSGCM5R) StoreBits(mask, b CSGCM5R_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCM5R) SetBits(mask CSGCM5R_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CSGCM5R) ClearBits(mask CSGCM5R_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCM5R) Load() CSGCM5R_Bits                  { return CSGCM5R_Bits(r.U32.Load()) }
func (r *CSGCM5R) Store(b CSGCM5R_Bits)                { r.U32.Store(uint32(b)) }

type CSGCM5R_Mask struct{ mmio.UM32 }

func (rm CSGCM5R_Mask) Load() CSGCM5R_Bits   { return CSGCM5R_Bits(rm.UM32.Load()) }
func (rm CSGCM5R_Mask) Store(b CSGCM5R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCM6R_Bits uint32

func (b CSGCM6R_Bits) Field(mask CSGCM6R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCM6R_Bits) J(v int) CSGCM6R_Bits {
	return CSGCM6R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCM6R struct{ mmio.U32 }

func (r *CSGCM6R) Bits(mask CSGCM6R_Bits) CSGCM6R_Bits { return CSGCM6R_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSGCM6R) StoreBits(mask, b CSGCM6R_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCM6R) SetBits(mask CSGCM6R_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CSGCM6R) ClearBits(mask CSGCM6R_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCM6R) Load() CSGCM6R_Bits                  { return CSGCM6R_Bits(r.U32.Load()) }
func (r *CSGCM6R) Store(b CSGCM6R_Bits)                { r.U32.Store(uint32(b)) }

type CSGCM6R_Mask struct{ mmio.UM32 }

func (rm CSGCM6R_Mask) Load() CSGCM6R_Bits   { return CSGCM6R_Bits(rm.UM32.Load()) }
func (rm CSGCM6R_Mask) Store(b CSGCM6R_Bits) { rm.UM32.Store(uint32(b)) }

type CSGCM7R_Bits uint32

func (b CSGCM7R_Bits) Field(mask CSGCM7R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSGCM7R_Bits) J(v int) CSGCM7R_Bits {
	return CSGCM7R_Bits(bits.Make32(v, uint32(mask)))
}

type CSGCM7R struct{ mmio.U32 }

func (r *CSGCM7R) Bits(mask CSGCM7R_Bits) CSGCM7R_Bits { return CSGCM7R_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSGCM7R) StoreBits(mask, b CSGCM7R_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSGCM7R) SetBits(mask CSGCM7R_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CSGCM7R) ClearBits(mask CSGCM7R_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CSGCM7R) Load() CSGCM7R_Bits                  { return CSGCM7R_Bits(r.U32.Load()) }
func (r *CSGCM7R) Store(b CSGCM7R_Bits)                { r.U32.Store(uint32(b)) }

type CSGCM7R_Mask struct{ mmio.UM32 }

func (rm CSGCM7R_Mask) Load() CSGCM7R_Bits   { return CSGCM7R_Bits(rm.UM32.Load()) }
func (rm CSGCM7R_Mask) Store(b CSGCM7R_Bits) { rm.UM32.Store(uint32(b)) }
