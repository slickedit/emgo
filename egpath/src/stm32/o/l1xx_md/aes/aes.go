// Peripheral: AES_Periph  AES hardware accelerator.
// Instances:
//  AES  mmap.AES_BASE
// Registers:
//  0x00 32  CR    Control register.
//  0x04 32  SR    Status register.
//  0x08 32  DINR  Data input register.
//  0x0C 32  DOUTR Data output register.
//  0x10 32  KEYR0 Key register 0.
//  0x14 32  KEYR1 Key register 1.
//  0x18 32  KEYR2 Key register 2.
//  0x1C 32  KEYR3 Key register 3.
//  0x20 32  IVR0  Initialization vector register 0.
//  0x24 32  IVR1  Initialization vector register 1.
//  0x28 32  IVR2  Initialization vector register 2.
//  0x2C 32  IVR3  Initialization vector register 3.
// Import:
//  stm32/o/l1xx_md/mmap
package aes

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	EN         CR_Bits = 0x01 << 0  //+ AES Enable.
	DATATYPE   CR_Bits = 0x03 << 1  //+ Data type selection.
	DATATYPE_0 CR_Bits = 0x01 << 1  //  Bit 0.
	DATATYPE_1 CR_Bits = 0x02 << 1  //  Bit 1.
	MODE       CR_Bits = 0x03 << 3  //+ AES Mode Of Operation.
	MODE_0     CR_Bits = 0x01 << 3  //  Bit 0.
	MODE_1     CR_Bits = 0x02 << 3  //  Bit 1.
	CHMOD      CR_Bits = 0x03 << 5  //+ AES Chaining Mode.
	CHMOD_0    CR_Bits = 0x01 << 5  //  Bit 0.
	CHMOD_1    CR_Bits = 0x02 << 5  //  Bit 1.
	CCFC       CR_Bits = 0x01 << 7  //+ Computation Complete Flag Clear.
	ERRC       CR_Bits = 0x01 << 8  //+ Error Clear.
	CCIE       CR_Bits = 0x01 << 9  //+ Computation Complete Interrupt Enable.
	ERRIE      CR_Bits = 0x01 << 10 //+ Error Interrupt Enable.
	DMAINEN    CR_Bits = 0x01 << 11 //+ DMA ENable managing the data input phase.
	DMAOUTEN   CR_Bits = 0x01 << 12 //+ DMA Enable managing the data output phase.
)

const (
	CCF   SR_Bits = 0x01 << 0 //+ Computation Complete Flag.
	RDERR SR_Bits = 0x01 << 1 //+ Read Error Flag.
	WRERR SR_Bits = 0x01 << 2 //+ Write Error Flag.
)