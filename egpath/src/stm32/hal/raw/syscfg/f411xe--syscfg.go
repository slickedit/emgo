// +build f411xe

// Peripheral: SYSCFG_Periph  System configuration controller.
// Instances:
//  SYSCFG  mmap.SYSCFG_BASE
// Registers:
//  0x00 32  MEMRMP    Memory remap register.
//  0x04 32  PMC       Peripheral mode configuration register.
//  0x08 32  EXTICR[4] External interrupt configuration registers.
//  0x20 32  CMPCR     Compensation cell control register.
// Import:
//  stm32/o/f411xe/mmap
package syscfg

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MEM_MODE   MEMRMP_Bits = 0x07 << 0  //+ SYSCFG_Memory Remap Config.
	MEM_MODE_0 MEMRMP_Bits = 0x01 << 0  //  Bit 0.
	MEM_MODE_1 MEMRMP_Bits = 0x02 << 0  //  Bit 1.
	MEM_MODE_2 MEMRMP_Bits = 0x04 << 0  //  Bit 2.
	FB_MODE    MEMRMP_Bits = 0x01 << 8  //+ User Flash Bank mode.
	SWP_FMC    MEMRMP_Bits = 0x03 << 10 //+ FMC memory mapping swap.
	SWP_FMC_0  MEMRMP_Bits = 0x01 << 10 //  Bit 0.
	SWP_FMC_1  MEMRMP_Bits = 0x02 << 10 //  Bit 1.
)

const (
	ADCxDC2      PMC_Bits = 0x07 << 16 //+ Refer to AN4073 on how to use this bit.
	ADC1DC2      PMC_Bits = 0x01 << 16 //  Refer to AN4073 on how to use this bit.
	ADC2DC2      PMC_Bits = 0x02 << 16 //  Refer to AN4073 on how to use this bit.
	ADC3DC2      PMC_Bits = 0x04 << 16 //  Refer to AN4073 on how to use this bit.
	MII_RMII_SEL PMC_Bits = 0x01 << 23 //+ Ethernet PHY interface selection.
)

const (
	CMP_PD CMPCR_Bits = 0x01 << 0 //+ Compensation cell ready flag.
	READY  CMPCR_Bits = 0x01 << 8 //+ Compensation cell power-down.
)
