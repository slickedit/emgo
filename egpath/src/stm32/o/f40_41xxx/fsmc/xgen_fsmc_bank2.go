package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type FSMC_Bank2_Periph struct {
	PCR2  PCR2
	SR2   SR2
	PMEM2 PMEM2
	PATT2 PATT2
	_     uint32
	ECCR2 ECCR2
}

func (p *FSMC_Bank2_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FSMC_Bank2 = (*FSMC_Bank2_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank2_R_BASE)))

type PCR2_Bits uint32

func (b PCR2_Bits) Field(mask PCR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCR2_Bits) J(v int) PCR2_Bits {
	return PCR2_Bits(bits.Make32(v, uint32(mask)))
}

type PCR2 struct{ mmio.U32 }

func (r *PCR2) Bits(mask PCR2_Bits) PCR2_Bits { return PCR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *PCR2) StoreBits(mask, b PCR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PCR2) SetBits(mask PCR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PCR2) ClearBits(mask PCR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PCR2) Load() PCR2_Bits               { return PCR2_Bits(r.U32.Load()) }
func (r *PCR2) Store(b PCR2_Bits)             { r.U32.Store(uint32(b)) }

type PCR2_Mask struct{ mmio.UM32 }

func (rm PCR2_Mask) Load() PCR2_Bits   { return PCR2_Bits(rm.UM32.Load()) }
func (rm PCR2_Mask) Store(b PCR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank2_Periph) PWAITEN() PCR2_Mask {
	return PCR2_Mask{mmio.UM32{&p.PCR2.U32, uint32(PWAITEN)}}
}

func (p *FSMC_Bank2_Periph) PBKEN() PCR2_Mask {
	return PCR2_Mask{mmio.UM32{&p.PCR2.U32, uint32(PBKEN)}}
}

func (p *FSMC_Bank2_Periph) PTYP() PCR2_Mask {
	return PCR2_Mask{mmio.UM32{&p.PCR2.U32, uint32(PTYP)}}
}

func (p *FSMC_Bank2_Periph) PWID() PCR2_Mask {
	return PCR2_Mask{mmio.UM32{&p.PCR2.U32, uint32(PWID)}}
}

func (p *FSMC_Bank2_Periph) ECCEN() PCR2_Mask {
	return PCR2_Mask{mmio.UM32{&p.PCR2.U32, uint32(ECCEN)}}
}

func (p *FSMC_Bank2_Periph) TCLR() PCR2_Mask {
	return PCR2_Mask{mmio.UM32{&p.PCR2.U32, uint32(TCLR)}}
}

func (p *FSMC_Bank2_Periph) TAR() PCR2_Mask {
	return PCR2_Mask{mmio.UM32{&p.PCR2.U32, uint32(TAR)}}
}

func (p *FSMC_Bank2_Periph) ECCPS() PCR2_Mask {
	return PCR2_Mask{mmio.UM32{&p.PCR2.U32, uint32(ECCPS)}}
}

type SR2_Bits uint32

func (b SR2_Bits) Field(mask SR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR2_Bits) J(v int) SR2_Bits {
	return SR2_Bits(bits.Make32(v, uint32(mask)))
}

type SR2 struct{ mmio.U32 }

func (r *SR2) Bits(mask SR2_Bits) SR2_Bits { return SR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR2) StoreBits(mask, b SR2_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR2) SetBits(mask SR2_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *SR2) ClearBits(mask SR2_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *SR2) Load() SR2_Bits              { return SR2_Bits(r.U32.Load()) }
func (r *SR2) Store(b SR2_Bits)            { r.U32.Store(uint32(b)) }

type SR2_Mask struct{ mmio.UM32 }

func (rm SR2_Mask) Load() SR2_Bits   { return SR2_Bits(rm.UM32.Load()) }
func (rm SR2_Mask) Store(b SR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank2_Periph) IRS() SR2_Mask {
	return SR2_Mask{mmio.UM32{&p.SR2.U32, uint32(IRS)}}
}

func (p *FSMC_Bank2_Periph) ILS() SR2_Mask {
	return SR2_Mask{mmio.UM32{&p.SR2.U32, uint32(ILS)}}
}

func (p *FSMC_Bank2_Periph) IFS() SR2_Mask {
	return SR2_Mask{mmio.UM32{&p.SR2.U32, uint32(IFS)}}
}

func (p *FSMC_Bank2_Periph) IREN() SR2_Mask {
	return SR2_Mask{mmio.UM32{&p.SR2.U32, uint32(IREN)}}
}

func (p *FSMC_Bank2_Periph) ILEN() SR2_Mask {
	return SR2_Mask{mmio.UM32{&p.SR2.U32, uint32(ILEN)}}
}

func (p *FSMC_Bank2_Periph) IFEN() SR2_Mask {
	return SR2_Mask{mmio.UM32{&p.SR2.U32, uint32(IFEN)}}
}

func (p *FSMC_Bank2_Periph) FEMPT() SR2_Mask {
	return SR2_Mask{mmio.UM32{&p.SR2.U32, uint32(FEMPT)}}
}

type PMEM2_Bits uint32

func (b PMEM2_Bits) Field(mask PMEM2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PMEM2_Bits) J(v int) PMEM2_Bits {
	return PMEM2_Bits(bits.Make32(v, uint32(mask)))
}

type PMEM2 struct{ mmio.U32 }

func (r *PMEM2) Bits(mask PMEM2_Bits) PMEM2_Bits { return PMEM2_Bits(r.U32.Bits(uint32(mask))) }
func (r *PMEM2) StoreBits(mask, b PMEM2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PMEM2) SetBits(mask PMEM2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PMEM2) ClearBits(mask PMEM2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PMEM2) Load() PMEM2_Bits                { return PMEM2_Bits(r.U32.Load()) }
func (r *PMEM2) Store(b PMEM2_Bits)              { r.U32.Store(uint32(b)) }

type PMEM2_Mask struct{ mmio.UM32 }

func (rm PMEM2_Mask) Load() PMEM2_Bits   { return PMEM2_Bits(rm.UM32.Load()) }
func (rm PMEM2_Mask) Store(b PMEM2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank2_Periph) MEMSET2() PMEM2_Mask {
	return PMEM2_Mask{mmio.UM32{&p.PMEM2.U32, uint32(MEMSET2)}}
}

func (p *FSMC_Bank2_Periph) MEMWAIT2() PMEM2_Mask {
	return PMEM2_Mask{mmio.UM32{&p.PMEM2.U32, uint32(MEMWAIT2)}}
}

func (p *FSMC_Bank2_Periph) MEMHOLD2() PMEM2_Mask {
	return PMEM2_Mask{mmio.UM32{&p.PMEM2.U32, uint32(MEMHOLD2)}}
}

func (p *FSMC_Bank2_Periph) MEMHIZ2() PMEM2_Mask {
	return PMEM2_Mask{mmio.UM32{&p.PMEM2.U32, uint32(MEMHIZ2)}}
}

type PATT2_Bits uint32

func (b PATT2_Bits) Field(mask PATT2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PATT2_Bits) J(v int) PATT2_Bits {
	return PATT2_Bits(bits.Make32(v, uint32(mask)))
}

type PATT2 struct{ mmio.U32 }

func (r *PATT2) Bits(mask PATT2_Bits) PATT2_Bits { return PATT2_Bits(r.U32.Bits(uint32(mask))) }
func (r *PATT2) StoreBits(mask, b PATT2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PATT2) SetBits(mask PATT2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PATT2) ClearBits(mask PATT2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PATT2) Load() PATT2_Bits                { return PATT2_Bits(r.U32.Load()) }
func (r *PATT2) Store(b PATT2_Bits)              { r.U32.Store(uint32(b)) }

type PATT2_Mask struct{ mmio.UM32 }

func (rm PATT2_Mask) Load() PATT2_Bits   { return PATT2_Bits(rm.UM32.Load()) }
func (rm PATT2_Mask) Store(b PATT2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank2_Periph) ATTSET2() PATT2_Mask {
	return PATT2_Mask{mmio.UM32{&p.PATT2.U32, uint32(ATTSET2)}}
}

func (p *FSMC_Bank2_Periph) ATTWAIT2() PATT2_Mask {
	return PATT2_Mask{mmio.UM32{&p.PATT2.U32, uint32(ATTWAIT2)}}
}

func (p *FSMC_Bank2_Periph) ATTHOLD2() PATT2_Mask {
	return PATT2_Mask{mmio.UM32{&p.PATT2.U32, uint32(ATTHOLD2)}}
}

func (p *FSMC_Bank2_Periph) ATTHIZ2() PATT2_Mask {
	return PATT2_Mask{mmio.UM32{&p.PATT2.U32, uint32(ATTHIZ2)}}
}

type ECCR2_Bits uint32

func (b ECCR2_Bits) Field(mask ECCR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ECCR2_Bits) J(v int) ECCR2_Bits {
	return ECCR2_Bits(bits.Make32(v, uint32(mask)))
}

type ECCR2 struct{ mmio.U32 }

func (r *ECCR2) Bits(mask ECCR2_Bits) ECCR2_Bits { return ECCR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *ECCR2) StoreBits(mask, b ECCR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ECCR2) SetBits(mask ECCR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ECCR2) ClearBits(mask ECCR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ECCR2) Load() ECCR2_Bits                { return ECCR2_Bits(r.U32.Load()) }
func (r *ECCR2) Store(b ECCR2_Bits)              { r.U32.Store(uint32(b)) }

type ECCR2_Mask struct{ mmio.UM32 }

func (rm ECCR2_Mask) Load() ECCR2_Bits   { return ECCR2_Bits(rm.UM32.Load()) }
func (rm ECCR2_Mask) Store(b ECCR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank2_Periph) ECC2() ECCR2_Mask {
	return ECCR2_Mask{mmio.UM32{&p.ECCR2.U32, uint32(ECC2)}}
}
