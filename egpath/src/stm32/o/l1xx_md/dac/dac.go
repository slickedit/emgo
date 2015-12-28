// Peripheral: DAC_Periph  Digital to Analog Converter.
// Instances:
//  DAC  mmap.DAC_BASE
// Registers:
//  0x00 32  CR      Control register.
//  0x04 32  SWTRIGR Software trigger register.
//  0x08 32  DHR12R1 Channel1 12-bit right-aligned data holding register.
//  0x0C 32  DHR12L1 Channel1 12-bit left aligned data holding register.
//  0x10 32  DHR8R1  Channel1 8-bit right aligned data holding register.
//  0x14 32  DHR12R2 Channel2 12-bit right aligned data holding register.
//  0x18 32  DHR12L2 Channel2 12-bit left aligned data holding register.
//  0x1C 32  DHR8R2  Channel2 8-bit right-aligned data holding register.
//  0x20 32  DHR12RD Dual DAC 12-bit right-aligned data holding register.
//  0x24 32  DHR12LD DUAL DAC 12-bit left aligned data holding register.
//  0x28 32  DHR8RD  DUAL DAC 8-bit right aligned data holding register.
//  0x2C 32  DOR1    Channel1 data output register.
//  0x30 32  DOR2    Channel2 data output register.
//  0x34 32  SR      Status register.
// Import:
//  stm32/o/l1xx_md/mmap
package dac

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	EN1       CR_Bits = 0x01 << 0  //+ DAC channel1 enable.
	BOFF1     CR_Bits = 0x01 << 1  //+ DAC channel1 output buffer disable.
	TEN1      CR_Bits = 0x01 << 2  //+ DAC channel1 Trigger enable.
	TSEL1     CR_Bits = 0x07 << 3  //+ TSEL1[2:0] (DAC channel1 Trigger selection).
	TSEL1_0   CR_Bits = 0x01 << 3  //  Bit 0.
	TSEL1_1   CR_Bits = 0x02 << 3  //  Bit 1.
	TSEL1_2   CR_Bits = 0x04 << 3  //  Bit 2.
	WAVE1     CR_Bits = 0x03 << 6  //+ WAVE1[1:0] (DAC channel1 noise/triangle wave generation enable).
	WAVE1_0   CR_Bits = 0x01 << 6  //  Bit 0.
	WAVE1_1   CR_Bits = 0x02 << 6  //  Bit 1.
	MAMP1     CR_Bits = 0x0F << 8  //+ MAMP1[3:0] (DAC channel1 Mask/Amplitude selector).
	MAMP1_0   CR_Bits = 0x01 << 8  //  Bit 0.
	MAMP1_1   CR_Bits = 0x02 << 8  //  Bit 1.
	MAMP1_2   CR_Bits = 0x04 << 8  //  Bit 2.
	MAMP1_3   CR_Bits = 0x08 << 8  //  Bit 3.
	DMAEN1    CR_Bits = 0x01 << 12 //+ DAC channel1 DMA enable.
	DMAUDRIE1 CR_Bits = 0x01 << 13 //+ DAC channel1 DMA underrun interrupt enable.
	EN2       CR_Bits = 0x01 << 16 //+ DAC channel2 enable.
	BOFF2     CR_Bits = 0x01 << 17 //+ DAC channel2 output buffer disable.
	TEN2      CR_Bits = 0x01 << 18 //+ DAC channel2 Trigger enable.
	TSEL2     CR_Bits = 0x07 << 19 //+ TSEL2[2:0] (DAC channel2 Trigger selection).
	TSEL2_0   CR_Bits = 0x01 << 19 //  Bit 0.
	TSEL2_1   CR_Bits = 0x02 << 19 //  Bit 1.
	TSEL2_2   CR_Bits = 0x04 << 19 //  Bit 2.
	WAVE2     CR_Bits = 0x03 << 22 //+ WAVE2[1:0] (DAC channel2 noise/triangle wave generation enable).
	WAVE2_0   CR_Bits = 0x01 << 22 //  Bit 0.
	WAVE2_1   CR_Bits = 0x02 << 22 //  Bit 1.
	MAMP2     CR_Bits = 0x0F << 24 //+ MAMP2[3:0] (DAC channel2 Mask/Amplitude selector).
	MAMP2_0   CR_Bits = 0x01 << 24 //  Bit 0.
	MAMP2_1   CR_Bits = 0x02 << 24 //  Bit 1.
	MAMP2_2   CR_Bits = 0x04 << 24 //  Bit 2.
	MAMP2_3   CR_Bits = 0x08 << 24 //  Bit 3.
	DMAEN2    CR_Bits = 0x01 << 28 //+ DAC channel2 DMA enabled.
	DMAUDRIE2 CR_Bits = 0x01 << 29 //+ DAC channel2 DMA underrun interrupt enable.
)

const (
	SWTRIG1 SWTRIGR_Bits = 0x01 << 0 //+ DAC channel1 software trigger.
	SWTRIG2 SWTRIGR_Bits = 0x01 << 1 //+ DAC channel2 software trigger.
)

const (
	DACC1DHR DHR12R1_Bits = 0xFFF << 0 //+ DAC channel1 12-bit Right aligned data.
)

const (
	DACC1DHR DHR12L1_Bits = 0xFFF << 4 //+ DAC channel1 12-bit Left aligned data.
)

const (
	DACC1DHR DHR8R1_Bits = 0xFF << 0 //+ DAC channel1 8-bit Right aligned data.
)

const (
	DACC2DHR DHR12R2_Bits = 0xFFF << 0 //+ DAC channel2 12-bit Right aligned data.
)

const (
	DACC2DHR DHR12L2_Bits = 0xFFF << 4 //+ DAC channel2 12-bit Left aligned data.
)

const (
	DACC2DHR DHR8R2_Bits = 0xFF << 0 //+ DAC channel2 8-bit Right aligned data.
)

const (
	DACC1DHR DHR12RD_Bits = 0xFFF << 0  //+ DAC channel1 12-bit Right aligned data.
	DACC2DHR DHR12RD_Bits = 0xFFF << 16 //+ DAC channel2 12-bit Right aligned data.
)

const (
	DACC1DHR DHR12LD_Bits = 0xFFF << 4  //+ DAC channel1 12-bit Left aligned data.
	DACC2DHR DHR12LD_Bits = 0xFFF << 20 //+ DAC channel2 12-bit Left aligned data.
)

const (
	DACC1DHR DHR8RD_Bits = 0xFF << 0 //+ DAC channel1 8-bit Right aligned data.
	DACC2DHR DHR8RD_Bits = 0xFF << 8 //+ DAC channel2 8-bit Right aligned data.
)

const (
	DACC1DOR DOR1_Bits = 0xFFF << 0 //+ DAC channel1 data output.
)

const (
	DACC2DOR DOR2_Bits = 0xFFF << 0 //+ DAC channel2 data output.
)

const (
	DMAUDR1 SR_Bits = 0x01 << 13 //+ DAC channel1 DMA underrun flag.
	DMAUDR2 SR_Bits = 0x01 << 29 //+ DAC channel2 DMA underrun flag.
)
