package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func rightpad(s string, size int) []byte {
	b := make([]byte, size)

	// copy keeps old data, not overwritten, in dest
	copy(b, []byte(s)[:])

	// zero pad (make sure old data is gone)
	for i := len(s); i < size; i++ {
		b[i] = 0
	}

	return b
}

//  will add more at a later time

func (f *SPC_file) SetSongTitle(title string) {
	f.Song["song_title"] = rightpad(title, len(f.Song["song_title"]))
}

func (f *SPC_file) SetArtist(artist string) {
	f.Song["game_title"] = rightpad(artist, len(f.Song["game_title"]))
}

func (f *SPC_file) SetGameTitle(title string) {
	f.Song["game_title"] = rightpad(title, len(f.Song["game_title"]))
}

func (f *SPC_file) Save() error {
	/* default save as the 'Game Name - Song Name' */
	var filename string
	// must trim zero padding
	// fmt.Sprintf keeps the zero padding
	filename = fmt.Sprintf("%s - %s.spc",
		bytes.Trim(f.Song["game_title"], "\x00"),
		bytes.Trim(f.Song["song_title"], "\x00"))

	return f.SaveAs(filename)
}

func (f *SPC_file) SaveAs(filename string) error {
	buffer := make([]byte, 0)

	var counter uint

	for _, key := range header_keys {
		buffer = append(buffer, f.Headers[key]...)
		counter++
	}

	for _, key := range register_keys {
		buffer = append(buffer, f.Registers[key]...)
		counter++
	}

	for _, key := range metadata_keys {
		buffer = append(buffer, f.Song[key]...)
		counter++
	}

	for _, key := range ram_keys {
		buffer = append(buffer, f.Ram[key]...)
		counter++
	}
	buffer = append(buffer, f.Ram["extended_ID666"]...)

	fmt.Println("FILE SIZE:", len(buffer))
	return ioutil.WriteFile(filename, buffer, 0644)
}
