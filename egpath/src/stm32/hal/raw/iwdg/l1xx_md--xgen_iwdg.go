// +build l1xx_md

package iwdg

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type IWDG_Periph struct {
	KR  KR
	PR  PR
	RLR RLR
	SR  SR
}

func (p *IWDG_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var IWDG = (*IWDG_Periph)(unsafe.Pointer(uintptr(mmap.IWDG_BASE)))

type KR_Bits uint32

func (b KR_Bits) Field(mask KR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask KR_Bits) J(v int) KR_Bits {
	return KR_Bits(bits.Make32(v, uint32(mask)))
}

type KR struct{ mmio.U32 }

func (r *KR) Bits(mask KR_Bits) KR_Bits { return KR_Bits(r.U32.Bits(uint32(mask))) }
func (r *KR) StoreBits(mask, b KR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *KR) SetBits(mask KR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *KR) ClearBits(mask KR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *KR) Load() KR_Bits             { return KR_Bits(r.U32.Load()) }
func (r *KR) Store(b KR_Bits)           { r.U32.Store(uint32(b)) }

type KR_Mask struct{ mmio.UM32 }

func (rm KR_Mask) Load() KR_Bits   { return KR_Bits(rm.UM32.Load()) }
func (rm KR_Mask) Store(b KR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *IWDG_Periph) KEY() KR_Mask {
	return KR_Mask{mmio.UM32{&p.KR.U32, uint32(KEY)}}
}

type PR_Bits uint32

func (b PR_Bits) Field(mask PR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PR_Bits) J(v int) PR_Bits {
	return PR_Bits(bits.Make32(v, uint32(mask)))
}

type PR struct{ mmio.U32 }

func (r *PR) Bits(mask PR_Bits) PR_Bits { return PR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PR) StoreBits(mask, b PR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PR) SetBits(mask PR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *PR) ClearBits(mask PR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *PR) Load() PR_Bits             { return PR_Bits(r.U32.Load()) }
func (r *PR) Store(b PR_Bits)           { r.U32.Store(uint32(b)) }

type PR_Mask struct{ mmio.UM32 }

func (rm PR_Mask) Load() PR_Bits   { return PR_Bits(rm.UM32.Load()) }
func (rm PR_Mask) Store(b PR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *IWDG_Periph) PR() PR_Mask {
	return PR_Mask{mmio.UM32{&p.PR.U32, uint32(PR)}}
}

type RLR_Bits uint32

func (b RLR_Bits) Field(mask RLR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RLR_Bits) J(v int) RLR_Bits {
	return RLR_Bits(bits.Make32(v, uint32(mask)))
}

type RLR struct{ mmio.U32 }

func (r *RLR) Bits(mask RLR_Bits) RLR_Bits { return RLR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RLR) StoreBits(mask, b RLR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLR) SetBits(mask RLR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *RLR) ClearBits(mask RLR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *RLR) Load() RLR_Bits              { return RLR_Bits(r.U32.Load()) }
func (r *RLR) Store(b RLR_Bits)            { r.U32.Store(uint32(b)) }

type RLR_Mask struct{ mmio.UM32 }

func (rm RLR_Mask) Load() RLR_Bits   { return RLR_Bits(rm.UM32.Load()) }
func (rm RLR_Mask) Store(b RLR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *IWDG_Periph) RL() RLR_Mask {
	return RLR_Mask{mmio.UM32{&p.RLR.U32, uint32(RL)}}
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

func (p *IWDG_Periph) PVU() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(PVU)}}
}

func (p *IWDG_Periph) RVU() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(RVU)}}
}
