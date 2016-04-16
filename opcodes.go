package main

func dp_offset(psw byte) int {
	// direct page offset
	// if DP = 1 -> offset = 0x0100
	// else      -> offset = 0x0000
	if psw&DIRECTPAGE > 0 {
		return 0x0100
	}
	return 0x0000
}

var opcodes = []func(s *SPC700){
	// jumptable

	func(s *SPC700) { // 0x00
		// NOP
		s.PC += 2
	},

	func(s *SPC700) { // 0x01
		// CALL [$FFDE]
		s.PC += 8
	},

	func(s *SPC700) { // 0x02
		// d.0 = 1
		s.RAM[dp_offset(s.PSW)] = d0 // bitmask 00000001
		s.PC += 2
	},

	func(s *SPC700) { //0x03
		// PC+=r if d.0 == 1
		// r = relative offset 5/7
		if s.RAM[dp_offset(s.PSW)]&1 == 1 {
			s.PC += 7
		} else {
			s.PC += 5
		}
	},

	func(s *SPC700) { //0x04
		// OR A, direct page
		ram_offset := dp_offset(s.PSW)
		s.A |= s.RAM[ram_offset]
		s.PC += 3
	},

	func(s *SPC700) { // 0x05
		// A | (abs)
		s.PC += 4
	},

	func(s *SPC700) { // 0x06
		// OR A, (X)
		s.A |= s.X
		s.PC += 3
	},

	func(s *SPC700) { // 0x07
		// OR A, [dp+X]
		s.A |= s.RAM[s.RAM[dp_offset(s.PSW)]+s.X]
		s.PC += 6
	},
}

func (s *SPC700) ProcessOP(opcode byte) {
	opcodes[opcode](s)
}
