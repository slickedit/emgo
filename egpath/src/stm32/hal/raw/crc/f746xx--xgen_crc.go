// +build f746xx

package crc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type CRC_Periph struct {
	DR   DR
	IDR  IDR
	_    uint8
	_    uint16
	CR   CR
	_    uint32
	INIT INIT
	POL  POL
}

func (p *CRC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var CRC = (*CRC_Periph)(unsafe.Pointer(uintptr(mmap.CRC_BASE)))

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

type IDR_Bits uint8

func (b IDR_Bits) Field(mask IDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDR_Bits) J(v int) IDR_Bits {
	return IDR_Bits(bits.Make32(v, uint32(mask)))
}

type IDR struct{ mmio.U8 }

func (r *IDR) Bits(mask IDR_Bits) IDR_Bits { return IDR_Bits(r.U8.Bits(uint8(mask))) }
func (r *IDR) StoreBits(mask, b IDR_Bits)  { r.U8.StoreBits(uint8(mask), uint8(b)) }
func (r *IDR) SetBits(mask IDR_Bits)       { r.U8.SetBits(uint8(mask)) }
func (r *IDR) ClearBits(mask IDR_Bits)     { r.U8.ClearBits(uint8(mask)) }
func (r *IDR) Load() IDR_Bits              { return IDR_Bits(r.U8.Load()) }
func (r *IDR) Store(b IDR_Bits)            { r.U8.Store(uint8(b)) }

type IDR_Mask struct{ mmio.UM8 }

func (rm IDR_Mask) Load() IDR_Bits   { return IDR_Bits(rm.UM8.Load()) }
func (rm IDR_Mask) Store(b IDR_Bits) { rm.UM8.Store(uint8(b)) }

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

func (p *CRC_Periph) RESET() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(RESET)}}
}

func (p *CRC_Periph) POLYSIZE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(POLYSIZE)}}
}

func (p *CRC_Periph) REV_IN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(REV_IN)}}
}

func (p *CRC_Periph) REV_OUT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(REV_OUT)}}
}

type INIT_Bits uint32

func (b INIT_Bits) Field(mask INIT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask INIT_Bits) J(v int) INIT_Bits {
	return INIT_Bits(bits.Make32(v, uint32(mask)))
}

type INIT struct{ mmio.U32 }

func (r *INIT) Bits(mask INIT_Bits) INIT_Bits { return INIT_Bits(r.U32.Bits(uint32(mask))) }
func (r *INIT) StoreBits(mask, b INIT_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *INIT) SetBits(mask INIT_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *INIT) ClearBits(mask INIT_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *INIT) Load() INIT_Bits               { return INIT_Bits(r.U32.Load()) }
func (r *INIT) Store(b INIT_Bits)             { r.U32.Store(uint32(b)) }

type INIT_Mask struct{ mmio.UM32 }

func (rm INIT_Mask) Load() INIT_Bits   { return INIT_Bits(rm.UM32.Load()) }
func (rm INIT_Mask) Store(b INIT_Bits) { rm.UM32.Store(uint32(b)) }

type POL_Bits uint32

func (b POL_Bits) Field(mask POL_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask POL_Bits) J(v int) POL_Bits {
	return POL_Bits(bits.Make32(v, uint32(mask)))
}

type POL struct{ mmio.U32 }

func (r *POL) Bits(mask POL_Bits) POL_Bits { return POL_Bits(r.U32.Bits(uint32(mask))) }
func (r *POL) StoreBits(mask, b POL_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *POL) SetBits(mask POL_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *POL) ClearBits(mask POL_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *POL) Load() POL_Bits              { return POL_Bits(r.U32.Load()) }
func (r *POL) Store(b POL_Bits)            { r.U32.Store(uint32(b)) }

type POL_Mask struct{ mmio.UM32 }

func (rm POL_Mask) Load() POL_Bits   { return POL_Bits(rm.UM32.Load()) }
func (rm POL_Mask) Store(b POL_Bits) { rm.UM32.Store(uint32(b)) }
