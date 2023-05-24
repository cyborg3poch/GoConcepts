package main

import (
	"fmt"
)

type Song struct {
	songname, artist string
}

type Playlist struct {
	genre string
	songs []Song
}

func main() {

	songs := []Song{
		{songname: "Kal ho na ho ", artist: "Sonu Nigam"},
		{songname: "Tum hi ho ", artist: "Arijit SIngh"},
	}

	playlist := Playlist{genre: "Romantic", songs: songs}

	for _, s := range playlist.songs {
		fmt.Printf("%-20s %20s", s.artist, s.songname)
	}
}
