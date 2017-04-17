// +build l1xx_md

package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type FSMC_Bank1E_Periph struct {
	BWTR [7]BWTR
}

func (p *FSMC_Bank1E_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FSMC_Bank1E = (*FSMC_Bank1E_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank1E_R_BASE)))

type BWTR_Bits uint32

func (b BWTR_Bits) Field(mask BWTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BWTR_Bits) J(v int) BWTR_Bits {
	return BWTR_Bits(bits.Make32(v, uint32(mask)))
}

type BWTR struct{ mmio.U32 }

func (r *BWTR) Bits(mask BWTR_Bits) BWTR_Bits { return BWTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BWTR) StoreBits(mask, b BWTR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BWTR) SetBits(mask BWTR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BWTR) ClearBits(mask BWTR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BWTR) Load() BWTR_Bits               { return BWTR_Bits(r.U32.Load()) }
func (r *BWTR) Store(b BWTR_Bits)             { r.U32.Store(uint32(b)) }

type BWTR_Mask struct{ mmio.UM32 }

func (rm BWTR_Mask) Load() BWTR_Bits   { return BWTR_Bits(rm.UM32.Load()) }
func (rm BWTR_Mask) Store(b BWTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank1E_Periph) EADDSET(n int) BWTR_Mask {
	return BWTR_Mask{mmio.UM32{&p.BWTR[n].U32, uint32(EADDSET)}}
}

func (p *FSMC_Bank1E_Periph) EADDHLD(n int) BWTR_Mask {
	return BWTR_Mask{mmio.UM32{&p.BWTR[n].U32, uint32(EADDHLD)}}
}

func (p *FSMC_Bank1E_Periph) EDATAST(n int) BWTR_Mask {
	return BWTR_Mask{mmio.UM32{&p.BWTR[n].U32, uint32(EDATAST)}}
}

func (p *FSMC_Bank1E_Periph) ECLKDIV(n int) BWTR_Mask {
	return BWTR_Mask{mmio.UM32{&p.BWTR[n].U32, uint32(ECLKDIV)}}
}

func (p *FSMC_Bank1E_Periph) EDATLAT(n int) BWTR_Mask {
	return BWTR_Mask{mmio.UM32{&p.BWTR[n].U32, uint32(EDATLAT)}}
}

func (p *FSMC_Bank1E_Periph) EACCMOD(n int) BWTR_Mask {
	return BWTR_Mask{mmio.UM32{&p.BWTR[n].U32, uint32(EACCMOD)}}
}
