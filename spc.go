package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough args")
		return
	}
	filename := os.Args[1]

	s := NewSPC()
	s.Decode(filename)

	fmt.Println(s.Song)
	fmt.Println(string(s.Song["song_title"]))
	fmt.Println(string(s.Song["artist"]))
	fmt.Println(string(s.Song["dumper_name"]))
	fmt.Println(string(s.Song["game_title"]))

	s.SetSongTitle("Memes")
	s.SetGameTitle("Dank")
	fmt.Printf("Saving... %s - %s\n", s.Song["game_title"], s.Song["song_title"])
	if err := s.Save(); err != nil {
		fmt.Println(err)
	}
}
