package main

import (
	"fmt"
	"github.com/ProfOak/spc/spc700"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough args")
		return
	}
	filename := os.Args[1]

	s := spc700.NewSPC()
	s.Decode(filename)

	fmt.Println(s.Song)
	fmt.Println(string(s.Song["song_title"]))
	fmt.Println(string(s.Song["artist"]))
	fmt.Println(string(s.Song["dumper_name"]))
	fmt.Println(string(s.Song["game_title"]))

	fmt.Println("Saving...")
	if err := s.Save(); err != nil {
		panic(err)
	}
}
