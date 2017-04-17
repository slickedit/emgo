// +build l476xx

package lptim

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type LPTIM_Periph struct {
	ISR  ISR
	ICR  ICR
	IER  IER
	CFGR CFGR
	CR   CR
	CMP  CMP
	ARR  ARR
	CNT  CNT
	OR   OR
}

func (p *LPTIM_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var LPTIM1 = (*LPTIM_Periph)(unsafe.Pointer(uintptr(mmap.LPTIM1_BASE)))

//emgo:const
var LPTIM2 = (*LPTIM_Periph)(unsafe.Pointer(uintptr(mmap.LPTIM2_BASE)))

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

func (p *LPTIM_Periph) CMPM() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(CMPM)}}
}

func (p *LPTIM_Periph) ARRM() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ARRM)}}
}

func (p *LPTIM_Periph) EXTTRIG() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(EXTTRIG)}}
}

func (p *LPTIM_Periph) CMPOK() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(CMPOK)}}
}

func (p *LPTIM_Periph) ARROK() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ARROK)}}
}

func (p *LPTIM_Periph) UP() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(UP)}}
}

func (p *LPTIM_Periph) DOWN() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(DOWN)}}
}

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

func (p *LPTIM_Periph) CMPMCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(CMPMCF)}}
}

func (p *LPTIM_Periph) ARRMCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(ARRMCF)}}
}

func (p *LPTIM_Periph) EXTTRIGCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(EXTTRIGCF)}}
}

func (p *LPTIM_Periph) CMPOKCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(CMPOKCF)}}
}

func (p *LPTIM_Periph) ARROKCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(ARROKCF)}}
}

func (p *LPTIM_Periph) UPCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(UPCF)}}
}

func (p *LPTIM_Periph) DOWNCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(DOWNCF)}}
}

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

func (p *LPTIM_Periph) CMPMIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(CMPMIE)}}
}

func (p *LPTIM_Periph) ARRMIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(ARRMIE)}}
}

func (p *LPTIM_Periph) EXTTRIGIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(EXTTRIGIE)}}
}

func (p *LPTIM_Periph) CMPOKIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(CMPOKIE)}}
}

func (p *LPTIM_Periph) ARROKIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(ARROKIE)}}
}

func (p *LPTIM_Periph) UPIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(UPIE)}}
}

func (p *LPTIM_Periph) DOWNIE() IER_Mask {
	return IER_Mask{mmio.UM32{&p.IER.U32, uint32(DOWNIE)}}
}

type CFGR_Bits uint32

func (b CFGR_Bits) Field(mask CFGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR_Bits) J(v int) CFGR_Bits {
	return CFGR_Bits(bits.Make32(v, uint32(mask)))
}

type CFGR struct{ mmio.U32 }

func (r *CFGR) Bits(mask CFGR_Bits) CFGR_Bits { return CFGR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFGR) StoreBits(mask, b CFGR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFGR) SetBits(mask CFGR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CFGR) ClearBits(mask CFGR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CFGR) Load() CFGR_Bits               { return CFGR_Bits(r.U32.Load()) }
func (r *CFGR) Store(b CFGR_Bits)             { r.U32.Store(uint32(b)) }

type CFGR_Mask struct{ mmio.UM32 }

func (rm CFGR_Mask) Load() CFGR_Bits   { return CFGR_Bits(rm.UM32.Load()) }
func (rm CFGR_Mask) Store(b CFGR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LPTIM_Periph) CKSEL() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(CKSEL)}}
}

func (p *LPTIM_Periph) CKPOL() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(CKPOL)}}
}

func (p *LPTIM_Periph) CKFLT() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(CKFLT)}}
}

func (p *LPTIM_Periph) TRGFLT() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(TRGFLT)}}
}

func (p *LPTIM_Periph) PRESC() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PRESC)}}
}

func (p *LPTIM_Periph) TRIGSEL() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(TRIGSEL)}}
}

func (p *LPTIM_Periph) TRIGEN() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(TRIGEN)}}
}

func (p *LPTIM_Periph) TIMOUT() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(TIMOUT)}}
}

func (p *LPTIM_Periph) WAVE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(WAVE)}}
}

func (p *LPTIM_Periph) WAVPOL() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(WAVPOL)}}
}

func (p *LPTIM_Periph) PRELOAD() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(PRELOAD)}}
}

func (p *LPTIM_Periph) COUNTMODE() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(COUNTMODE)}}
}

func (p *LPTIM_Periph) ENC() CFGR_Mask {
	return CFGR_Mask{mmio.UM32{&p.CFGR.U32, uint32(ENC)}}
}

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

func (p *LPTIM_Periph) ENABLE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ENABLE)}}
}

func (p *LPTIM_Periph) SNGSTRT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(SNGSTRT)}}
}

func (p *LPTIM_Periph) CNTSTRT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(CNTSTRT)}}
}

type CMP_Bits uint32

func (b CMP_Bits) Field(mask CMP_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMP_Bits) J(v int) CMP_Bits {
	return CMP_Bits(bits.Make32(v, uint32(mask)))
}

type CMP struct{ mmio.U32 }

func (r *CMP) Bits(mask CMP_Bits) CMP_Bits { return CMP_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMP) StoreBits(mask, b CMP_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMP) SetBits(mask CMP_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CMP) ClearBits(mask CMP_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CMP) Load() CMP_Bits              { return CMP_Bits(r.U32.Load()) }
func (r *CMP) Store(b CMP_Bits)            { r.U32.Store(uint32(b)) }

type CMP_Mask struct{ mmio.UM32 }

func (rm CMP_Mask) Load() CMP_Bits   { return CMP_Bits(rm.UM32.Load()) }
func (rm CMP_Mask) Store(b CMP_Bits) { rm.UM32.Store(uint32(b)) }

type ARR_Bits uint32

func (b ARR_Bits) Field(mask ARR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ARR_Bits) J(v int) ARR_Bits {
	return ARR_Bits(bits.Make32(v, uint32(mask)))
}

type ARR struct{ mmio.U32 }

func (r *ARR) Bits(mask ARR_Bits) ARR_Bits { return ARR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ARR) StoreBits(mask, b ARR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ARR) SetBits(mask ARR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ARR) ClearBits(mask ARR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ARR) Load() ARR_Bits              { return ARR_Bits(r.U32.Load()) }
func (r *ARR) Store(b ARR_Bits)            { r.U32.Store(uint32(b)) }

type ARR_Mask struct{ mmio.UM32 }

func (rm ARR_Mask) Load() ARR_Bits   { return ARR_Bits(rm.UM32.Load()) }
func (rm ARR_Mask) Store(b ARR_Bits) { rm.UM32.Store(uint32(b)) }

type CNT_Bits uint32

func (b CNT_Bits) Field(mask CNT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CNT_Bits) J(v int) CNT_Bits {
	return CNT_Bits(bits.Make32(v, uint32(mask)))
}

type CNT struct{ mmio.U32 }

func (r *CNT) Bits(mask CNT_Bits) CNT_Bits { return CNT_Bits(r.U32.Bits(uint32(mask))) }
func (r *CNT) StoreBits(mask, b CNT_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CNT) SetBits(mask CNT_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CNT) ClearBits(mask CNT_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CNT) Load() CNT_Bits              { return CNT_Bits(r.U32.Load()) }
func (r *CNT) Store(b CNT_Bits)            { r.U32.Store(uint32(b)) }

type CNT_Mask struct{ mmio.UM32 }

func (rm CNT_Mask) Load() CNT_Bits   { return CNT_Bits(rm.UM32.Load()) }
func (rm CNT_Mask) Store(b CNT_Bits) { rm.UM32.Store(uint32(b)) }

type OR_Bits uint32

func (b OR_Bits) Field(mask OR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OR_Bits) J(v int) OR_Bits {
	return OR_Bits(bits.Make32(v, uint32(mask)))
}

type OR struct{ mmio.U32 }

func (r *OR) Bits(mask OR_Bits) OR_Bits { return OR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OR) StoreBits(mask, b OR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OR) SetBits(mask OR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *OR) ClearBits(mask OR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *OR) Load() OR_Bits             { return OR_Bits(r.U32.Load()) }
func (r *OR) Store(b OR_Bits)           { r.U32.Store(uint32(b)) }

type OR_Mask struct{ mmio.UM32 }

func (rm OR_Mask) Load() OR_Bits   { return OR_Bits(rm.UM32.Load()) }
func (rm OR_Mask) Store(b OR_Bits) { rm.UM32.Store(uint32(b)) }
