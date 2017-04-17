package fmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type FMC_Bank1_Periph struct {
	BTCR [4]BTCR
}

func (p *FMC_Bank1_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FMC_Bank1_R = (*FMC_Bank1_Periph)(unsafe.Pointer(uintptr(mmap.FMC_Bank1_R_BASE)))

type BTCR struct {
	BCR BCR
	BTR BTR
}

type BCR_Bits uint32

func (b BCR_Bits) Field(mask BCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BCR_Bits) J(v int) BCR_Bits {
	return BCR_Bits(bits.Make32(v, uint32(mask)))
}

type BCR struct{ mmio.U32 }

func (r *BCR) Bits(mask BCR_Bits) BCR_Bits { return BCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BCR) StoreBits(mask, b BCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BCR) SetBits(mask BCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *BCR) ClearBits(mask BCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *BCR) Load() BCR_Bits              { return BCR_Bits(r.U32.Load()) }
func (r *BCR) Store(b BCR_Bits)            { r.U32.Store(uint32(b)) }

type BCR_Mask struct{ mmio.UM32 }

func (rm BCR_Mask) Load() BCR_Bits   { return BCR_Bits(rm.UM32.Load()) }
func (rm BCR_Mask) Store(b BCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FMC_Bank1_Periph) CCLKEN(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(CCLKEN)}}
}

type BTR_Bits uint32

func (b BTR_Bits) Field(mask BTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BTR_Bits) J(v int) BTR_Bits {
	return BTR_Bits(bits.Make32(v, uint32(mask)))
}

type BTR struct{ mmio.U32 }

func (r *BTR) Bits(mask BTR_Bits) BTR_Bits { return BTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BTR) StoreBits(mask, b BTR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BTR) SetBits(mask BTR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *BTR) ClearBits(mask BTR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *BTR) Load() BTR_Bits              { return BTR_Bits(r.U32.Load()) }
func (r *BTR) Store(b BTR_Bits)            { r.U32.Store(uint32(b)) }

type BTR_Mask struct{ mmio.UM32 }

func (rm BTR_Mask) Load() BTR_Bits   { return BTR_Bits(rm.UM32.Load()) }
func (rm BTR_Mask) Store(b BTR_Bits) { rm.UM32.Store(uint32(b)) }
