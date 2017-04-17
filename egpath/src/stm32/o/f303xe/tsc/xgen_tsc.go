package tsc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type TSC_Periph struct {
	CR     CR
	IER    IER
	ICR    ICR
	ISR    ISR
	IOHCR  IOHCR
	_      uint32
	IOASCR IOASCR
	_      uint32
	IOSCR  IOSCR
	_      uint32
	IOCCR  IOCCR
	_      uint32
	IOGCSR IOGCSR
	IOGXCR [8]IOGXCR
}

func (p *TSC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var TSC = (*TSC_Periph)(unsafe.Pointer(uintptr(mmap.TSC_BASE)))

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

type IER_Bits uint32

func (b IER_Bits) Field(mask IER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IER_Bits) J(v int) IER_Bits {
	return IER_Bits(bits.Make32(v, uint32(mask)))
}

type IER struct{ mmio.U32 }

func (r *IER) Bits(mask IER_Bits) IER_Bits { return IER_Bits(r.U32.Bits(uint32(mask))) }
func (r *IER) StoreBits(mask, b IER_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IER) SetBits(mask IER_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IER) ClearBits(mask IER_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IER) Load() IER_Bits              { return IER_Bits(r.U32.Load()) }
func (r *IER) Store(b IER_Bits)            { r.U32.Store(uint32(b)) }

type IER_Mask struct{ mmio.UM32 }

func (rm IER_Mask) Load() IER_Bits   { return IER_Bits(rm.UM32.Load()) }
func (rm IER_Mask) Store(b IER_Bits) { rm.UM32.Store(uint32(b)) }

type ICR_Bits uint32

func (b ICR_Bits) Field(mask ICR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICR_Bits) J(v int) ICR_Bits {
	return ICR_Bits(bits.Make32(v, uint32(mask)))
}

type ICR struct{ mmio.U32 }

func (r *ICR) Bits(mask ICR_Bits) ICR_Bits { return ICR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ICR) StoreBits(mask, b ICR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ICR) SetBits(mask ICR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ICR) ClearBits(mask ICR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ICR) Load() ICR_Bits              { return ICR_Bits(r.U32.Load()) }
func (r *ICR) Store(b ICR_Bits)            { r.U32.Store(uint32(b)) }

type ICR_Mask struct{ mmio.UM32 }

func (rm ICR_Mask) Load() ICR_Bits   { return ICR_Bits(rm.UM32.Load()) }
func (rm ICR_Mask) Store(b ICR_Bits) { rm.UM32.Store(uint32(b)) }

type ISR_Bits uint32

func (b ISR_Bits) Field(mask ISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ISR_Bits) J(v int) ISR_Bits {
	return ISR_Bits(bits.Make32(v, uint32(mask)))
}

type ISR struct{ mmio.U32 }

func (r *ISR) Bits(mask ISR_Bits) ISR_Bits { return ISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ISR) StoreBits(mask, b ISR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ISR) SetBits(mask ISR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ISR) ClearBits(mask ISR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ISR) Load() ISR_Bits              { return ISR_Bits(r.U32.Load()) }
func (r *ISR) Store(b ISR_Bits)            { r.U32.Store(uint32(b)) }

type ISR_Mask struct{ mmio.UM32 }

func (rm ISR_Mask) Load() ISR_Bits   { return ISR_Bits(rm.UM32.Load()) }
func (rm ISR_Mask) Store(b ISR_Bits) { rm.UM32.Store(uint32(b)) }

type IOHCR_Bits uint32

func (b IOHCR_Bits) Field(mask IOHCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IOHCR_Bits) J(v int) IOHCR_Bits {
	return IOHCR_Bits(bits.Make32(v, uint32(mask)))
}

type IOHCR struct{ mmio.U32 }

func (r *IOHCR) Bits(mask IOHCR_Bits) IOHCR_Bits { return IOHCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IOHCR) StoreBits(mask, b IOHCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IOHCR) SetBits(mask IOHCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *IOHCR) ClearBits(mask IOHCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *IOHCR) Load() IOHCR_Bits                { return IOHCR_Bits(r.U32.Load()) }
func (r *IOHCR) Store(b IOHCR_Bits)              { r.U32.Store(uint32(b)) }

type IOHCR_Mask struct{ mmio.UM32 }

func (rm IOHCR_Mask) Load() IOHCR_Bits   { return IOHCR_Bits(rm.UM32.Load()) }
func (rm IOHCR_Mask) Store(b IOHCR_Bits) { rm.UM32.Store(uint32(b)) }

type IOASCR_Bits uint32

func (b IOASCR_Bits) Field(mask IOASCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IOASCR_Bits) J(v int) IOASCR_Bits {
	return IOASCR_Bits(bits.Make32(v, uint32(mask)))
}

type IOASCR struct{ mmio.U32 }

func (r *IOASCR) Bits(mask IOASCR_Bits) IOASCR_Bits { return IOASCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IOASCR) StoreBits(mask, b IOASCR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IOASCR) SetBits(mask IOASCR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *IOASCR) ClearBits(mask IOASCR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *IOASCR) Load() IOASCR_Bits                 { return IOASCR_Bits(r.U32.Load()) }
func (r *IOASCR) Store(b IOASCR_Bits)               { r.U32.Store(uint32(b)) }

type IOASCR_Mask struct{ mmio.UM32 }

func (rm IOASCR_Mask) Load() IOASCR_Bits   { return IOASCR_Bits(rm.UM32.Load()) }
func (rm IOASCR_Mask) Store(b IOASCR_Bits) { rm.UM32.Store(uint32(b)) }

type IOSCR_Bits uint32

func (b IOSCR_Bits) Field(mask IOSCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IOSCR_Bits) J(v int) IOSCR_Bits {
	return IOSCR_Bits(bits.Make32(v, uint32(mask)))
}

type IOSCR struct{ mmio.U32 }

func (r *IOSCR) Bits(mask IOSCR_Bits) IOSCR_Bits { return IOSCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IOSCR) StoreBits(mask, b IOSCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IOSCR) SetBits(mask IOSCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *IOSCR) ClearBits(mask IOSCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *IOSCR) Load() IOSCR_Bits                { return IOSCR_Bits(r.U32.Load()) }
func (r *IOSCR) Store(b IOSCR_Bits)              { r.U32.Store(uint32(b)) }

type IOSCR_Mask struct{ mmio.UM32 }

func (rm IOSCR_Mask) Load() IOSCR_Bits   { return IOSCR_Bits(rm.UM32.Load()) }
func (rm IOSCR_Mask) Store(b IOSCR_Bits) { rm.UM32.Store(uint32(b)) }

type IOCCR_Bits uint32

func (b IOCCR_Bits) Field(mask IOCCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IOCCR_Bits) J(v int) IOCCR_Bits {
	return IOCCR_Bits(bits.Make32(v, uint32(mask)))
}

type IOCCR struct{ mmio.U32 }

func (r *IOCCR) Bits(mask IOCCR_Bits) IOCCR_Bits { return IOCCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IOCCR) StoreBits(mask, b IOCCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IOCCR) SetBits(mask IOCCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *IOCCR) ClearBits(mask IOCCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *IOCCR) Load() IOCCR_Bits                { return IOCCR_Bits(r.U32.Load()) }
func (r *IOCCR) Store(b IOCCR_Bits)              { r.U32.Store(uint32(b)) }

type IOCCR_Mask struct{ mmio.UM32 }

func (rm IOCCR_Mask) Load() IOCCR_Bits   { return IOCCR_Bits(rm.UM32.Load()) }
func (rm IOCCR_Mask) Store(b IOCCR_Bits) { rm.UM32.Store(uint32(b)) }

type IOGCSR_Bits uint32

func (b IOGCSR_Bits) Field(mask IOGCSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IOGCSR_Bits) J(v int) IOGCSR_Bits {
	return IOGCSR_Bits(bits.Make32(v, uint32(mask)))
}

type IOGCSR struct{ mmio.U32 }

func (r *IOGCSR) Bits(mask IOGCSR_Bits) IOGCSR_Bits { return IOGCSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IOGCSR) StoreBits(mask, b IOGCSR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IOGCSR) SetBits(mask IOGCSR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *IOGCSR) ClearBits(mask IOGCSR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *IOGCSR) Load() IOGCSR_Bits                 { return IOGCSR_Bits(r.U32.Load()) }
func (r *IOGCSR) Store(b IOGCSR_Bits)               { r.U32.Store(uint32(b)) }

type IOGCSR_Mask struct{ mmio.UM32 }

func (rm IOGCSR_Mask) Load() IOGCSR_Bits   { return IOGCSR_Bits(rm.UM32.Load()) }
func (rm IOGCSR_Mask) Store(b IOGCSR_Bits) { rm.UM32.Store(uint32(b)) }

type IOGXCR_Bits uint32

func (b IOGXCR_Bits) Field(mask IOGXCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IOGXCR_Bits) J(v int) IOGXCR_Bits {
	return IOGXCR_Bits(bits.Make32(v, uint32(mask)))
}

type IOGXCR struct{ mmio.U32 }

func (r *IOGXCR) Bits(mask IOGXCR_Bits) IOGXCR_Bits { return IOGXCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IOGXCR) StoreBits(mask, b IOGXCR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IOGXCR) SetBits(mask IOGXCR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *IOGXCR) ClearBits(mask IOGXCR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *IOGXCR) Load() IOGXCR_Bits                 { return IOGXCR_Bits(r.U32.Load()) }
func (r *IOGXCR) Store(b IOGXCR_Bits)               { r.U32.Store(uint32(b)) }

type IOGXCR_Mask struct{ mmio.UM32 }

func (rm IOGXCR_Mask) Load() IOGXCR_Bits   { return IOGXCR_Bits(rm.UM32.Load()) }
func (rm IOGXCR_Mask) Store(b IOGXCR_Bits) { rm.UM32.Store(uint32(b)) }
