package spc700

type SPC_file struct {
	Size      int
	Headers   map[string][]byte
	Registers map[string][]byte
	Song      map[string][]byte
	Ram       map[string][]byte
}

type SPC700_freeze struct {
	// Sound emulation data
	PC [2]byte // program counter
	A  byte    // accumulator
	X  byte    // index register
	Y  byte    // index register
	SP byte    // stack pointer
	// Program Status Word
	PSW byte // register contains various bits that effect operation of the CPU
	/*
		=== RAM LAYOUT ===

		0x0000 - 0x00EF PAGE 0
		0x00F0 - 0x00FF REGISTERS
		0x0100 - 0x01FF PAGE 1
		0x0200 - 0xFFBF MEMORY
		0xFFC0 - 0xFFFF MEMORY [READ/WRITE]
		0xFFC0 - 0xFFFF MEMORY [WRITE ONLY]
		0xFFC0 - 0xFFFF 64 BYTE IPL ROM [READ ONLY]
	*/
	RAM [0x10000]byte
	DSP [128]byte
}

type CPU struct {
	// CPU register set
	PC  [2]byte // program counter
	A   byte    // accumulator
	X   byte    // index register
	Y   byte    // index register
	SP  byte    // stack pointer
	PSW byte    // register contains various bits that effect operation of the CPU
}
