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
	PC  [2]byte
	A   byte
	X   byte
	Y   byte
	PSW byte
	SP  byte
	RAM [0x10000]byte
	DSP [128]byte
}
