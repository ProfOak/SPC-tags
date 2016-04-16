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
	var (
		f SPC_file
		//s SPC700
	)

	f = NewSPC()

	// Read file, order content
	f.Decode(filename)

	// Load content into SPC700 cpu registers
	// used for actual audio emulation
	//s = f.LoadCart()

	/* testing stuff */
	fmt.Println(f.Song)
	fmt.Println(string(f.Song["song_title"]))
	fmt.Println(string(f.Song["artist"]))
	fmt.Println(string(f.Song["dumper_name"]))
	fmt.Println(string(f.Song["game_title"]))

	f.SetSongTitle("Memes")
	f.SetGameTitle("Dank")
	fmt.Printf("Saving... %s - %s\n", f.Song["game_title"], f.Song["song_title"])
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}
	/* /testing stuff */
}
