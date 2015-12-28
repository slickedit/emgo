// +build l1xx_md

// Peripheral: SYSCFG_Periph  SysTem Configuration.
// Instances:
//  SYSCFG  mmap.SYSCFG_BASE
// Registers:
//  0x00 32  MEMRMP    Memory remap register.
//  0x04 32  PMC       Peripheral mode configuration register.
//  0x08 32  EXTICR[4] External interrupt configuration registers.
// Import:
//  stm32/o/l1xx_md/mmap
package syscfg

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MEM_MODE    MEMRMP_Bits = 0x03 << 0 //+ SYSCFG_Memory Remap Config.
	MEM_MODE_0  MEMRMP_Bits = 0x01 << 0 //  Bit 0.
	MEM_MODE_1  MEMRMP_Bits = 0x02 << 0 //  Bit 1.
	BOOT_MODE   MEMRMP_Bits = 0x03 << 8 //+ Boot mode Config.
	BOOT_MODE_0 MEMRMP_Bits = 0x01 << 8 //  Bit 0.
	BOOT_MODE_1 MEMRMP_Bits = 0x02 << 8 //  Bit 1.
)

const (
	USB_PU PMC_Bits = 0x01 << 0 //+ SYSCFG PMC.
)
