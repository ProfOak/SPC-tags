package main

type SPC_file map[string][]byte

type Registers struct {
	//Mneumonic    Desc            Control
	F0 byte // Undocumented         ?/W
	F1 byte // Control Register      /W
	F2 byte // DSP Register Address R/W
	F3 byte // DSP Register Data    R/W
	F4 byte // Port 0               R/W
	F5 byte // Port 1               R/W
	F6 byte // Port 2               R/W
	F7 byte // Port 3               R/W
	F8 byte // Regular Memory       R/W
	F9 byte // Regular Memory       R/W
	FA byte // Timer-0               /W
	FB byte // Timer-1               /W
	FC byte // Timer-2               /W
	FD byte // Counter-0            R/
	FE byte // Counter-1            R/
	FF byte // Counter-2            R/
}

type SPC700 struct {
	// Sound emulation data
	PC  uint16 // program counter
	A   byte   // accumulator
	X   byte   // index register
	Y   byte   // index register
	SP  byte   // stack pointer
	PSW byte   // Program Status Word (explained below)
	DSP [128]byte
	RAM [0x10000]byte
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
	REGISTER Registers
}

// === PSW layout ===
// bitmasks for PSW and direct page offsets
// [N, V, P, B, H, I, Z, C]
const NEGATIVE, d7 = 128, 128    // N (Negative)
const OVERFLOW, d6 = 64, 64      // V (Overflow)
const DIRECTPAGE, d5 = 32, 32    // P (Direct page)
const BREAK, d4 = 16, 16         // B (Break)
const HALFCARRY, d3 = 8, 8       // H (Half Carry)
const INTERUPTENABLED, d2 = 4, 4 // I (Interrupt enabled (unused))
const ZERO, d1 = 2, 2            // Z (Zero)
const CARRY, d0 = 1, 1           // C (Carry)
