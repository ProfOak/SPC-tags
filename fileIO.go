package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func rightpad(str string, size int) []byte {
	buffer := make([]byte, size)

	// copy keeps old data, not overwritten, in dest
	copy(buffer, []byte(str)[:])

	// zero pad (make sure old data is gone)
	for i := len(str); i < size; i++ {
		buffer[i] = 0
	}

	return buffer
}

//  will add more at a later time

func (f SPC_file) SetSongTitle(title string) {
	f["song_title"] = rightpad(title, len(f["song_title"]))
}

func (f SPC_file) SetArtist(artist string) {
	f["game_title"] = rightpad(artist, len(f["game_title"]))
}

func (f SPC_file) SetGameTitle(title string) {
	f["game_title"] = rightpad(title, len(f["game_title"]))
}

func (f SPC_file) Save() error {
	/* default save as the 'Game Name - Song Name' */
	var filename string
	// must trim zero padding
	// fmt.Sprintf keeps the zero padding
	filename = fmt.Sprintf("%s - %s.spc",
		bytes.Trim(f["game_title"], "\x00"),
		bytes.Trim(f["song_title"], "\x00"))

	return f.SaveAs(filename)
}

func (f SPC_file) SaveAs(filename string) error {
	buffer := make([]byte, 0)

	var counter uint

	for _, key := range keys {
		buffer = append(buffer, f[key]...)
		counter++
	}
	buffer = append(buffer, f["extended_ID666"]...)

	fmt.Println("FILE SIZE:", len(buffer))
	return ioutil.WriteFile(filename, buffer, 0644)
}
