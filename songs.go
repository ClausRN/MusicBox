package main

import (
	"path/filepath"
	"time"
	"fmt"
	"log"
)


type MusicLibrary struct {
	SongNo	int32
	CurrentTrackNo int32
	Songs	[]string
}

func NewSongList() (result MusicLibrary) {
	files, err := filepath.Glob("/music/*.mp3")
	if err != nil {
		log.Fatal(err)
	}
	result.SongNo = (int32)(len(files))
	result.Songs = files
	result.CurrentTrackNo = 0
	return
}

func (mylib *MusicLibrary) NextTrack() string {
        mylib.CurrentTrackNo++
	if (mylib.CurrentTrackNo >= mylib.SongNo) {
		mylib.CurrentTrackNo = 0
	}
	fmt.Printf("Time %v: Song selected: %s\n", time.Now(), mylib.Songs[mylib.CurrentTrackNo])
	return mylib.Songs[mylib.CurrentTrackNo]
}

