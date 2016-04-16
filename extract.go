package main

import (
	"io/ioutil"
)

var offsets = []int{
	0x00,    // File header “SNES-SPC700 Sound File Data v0.30”
	0x21,    // 0x26, 0x26
	0x23,    // 0x26 = Header Has ID666 Information / 0x27 = Header Has No ID666 Tag
	0x24,    // Version Minor (i.e. 30)
	0x25,    // PC
	0x27,    // A
	0x28,    // X
	0x29,    // Y
	0x2A,    // PSW
	0x2B,    // SP (Lower Byte)
	0x2C,    // Reserved
	0x2E,    // Song Title
	0x4E,    // Game Title
	0x6E,    // Name of Dumper
	0x7E,    // Comments
	0x9E,    // Date SPC was Dumped (MM/DD/YYYY)
	0xA9,    // Number of Seconds to Play Song before Fading Out
	0xAC,    // Length of Fade in Milliseconds
	0xB1,    // Artist of Song
	0xD1,    // Default Channel Disables (0=Enable, 1=Disable)
	0xD2,    // Emulator used to dump SPC (0=
	0xD3,    // Reserved(0x00)
	0x100,   // 64KB RAM
	0x10100, // DSP Registers
	0x10180, // Unused
	0x101C0, // Extra RAM (Memory Region used when the IPL ROM region is set to read-only)
	0x10200, // Extended ID666 starting point
}

var header_keys = []string{
	"header", "bits", "tags", "version_minor",
}

var register_keys = []string{
	"pc", "a", "x", "y", "psw", "dsp", "reserved",
}

var metadata_keys = []string{
	"song_title", "game_title", "dumper_name", "comments", "date_dumped",
	"num_of_sec_before_fade", "fade_length", "artist",
	"default_channel_disables", "emulator_used", "reserved",
}

var ram_keys = []string{
	"64k_ram", "dsp_registers", "unused", "extra_ram",
	// extended_ID666 handled outside of for loop
}

func chunk(f []byte, fr int, to int) []byte {
	return f[offsets[fr]:offsets[to]]
}

func NewSPC() SPC_file {
	var f SPC_file
	f.Headers = make(map[string][]byte)
	f.Registers = make(map[string][]byte)
	f.Song = make(map[string][]byte)
	f.Ram = make(map[string][]byte)
	return f
}

func (f *SPC_file) Decode(filename string) {
	contents, _ := ioutil.ReadFile(filename)
	var counter int

	for _, key := range header_keys {
		f.Headers[key] = chunk(contents, counter, counter+1)
		counter++
	}

	for _, key := range register_keys {
		f.Registers[key] = chunk(contents, counter, counter+1)
		counter++
	}

	for _, key := range metadata_keys {
		f.Song[key] = chunk(contents, counter, counter+1)
		counter++
	}

	for _, key := range ram_keys {
		f.Ram[key] = chunk(contents, counter, counter+1)
		counter++
	}
	f.Ram["extended_ID666"] = contents[offsets[counter]:]
}

func (f SPC_file) LoadCart() SPC700 {
	// temp variables for conversion
	var (
		pc  uint16 = uint16(f.Registers["pc"][0])<<8 + uint16(f.Registers["pc"][1])
		dsp [128]byte
		ram [0x10000]byte
	)

	copy(dsp[:], f.Registers["dsp"])
	copy(ram[:], f.Ram["64k_ram"])

	// correct conversions
	return SPC700{
		PC:  pc,
		A:   f.Registers["a"][0], // technically byte arrays
		X:   f.Registers["x"][0], // one byte long
		Y:   f.Registers["y"][0], // silly "conversions"
		SP:  0,
		PSW: f.Registers["psw"][0],
		DSP: dsp,
		RAM: ram,
	}
}
