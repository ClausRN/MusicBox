package main

import (
	"fmt"
	"time"
	"github.com/edudev/go-omx/omx"
)

type OMXPlayer struct {
	OMX 	*omx.OmxInterface
	Media	MusicLibrary
}

func NewOMXPlayer() (result OMXPlayer) {
	obj, warn := omx.NewOmxInterface()
	if warn != nil {
		fmt.Println("Error on OMX startup: ", warn)
		panic(warn)
	}
        result.OMX = obj
	result.Media = NewSongList()
	return
}

func (player *OMXPlayer)Stop() {
	if player.OMX.HasPlayer() {
		status, err := player.OMX.PlaybackStatus()
		if err != nil {
			fmt.Println("Stop error: ", err)
		}
		if status == "playing" {
			fmt.Printf("Time %v: Stop playing\n", time.Now())
			player.OMX.Pause()
			time.Sleep(time.Duration(250) * time.Millisecond)
		}
	}
}

func (player *OMXPlayer)Start() {
	if player.OMX.HasPlayer() == false {
		fmt.Printf("Time %v: Start playing\n", time.Now())
		player.OMX.StartPlayer(player.Media.NextTrack())
		time.Sleep(time.Duration(500) * time.Millisecond)
        } else {
		status, err := player.OMX.PlaybackStatus()
                if err != nil {
			fmt.Println("PlaybackStatus error: ", err)
                }
                if status != "playing" {
			fmt.Printf("Time %v: Restart playing\n", time.Now())
			player.OMX.Play()
			time.Sleep(time.Duration(250) * time.Millisecond)		
		} else {
		}
	}

}

func (player *OMXPlayer)Close() {
	player.OMX.Stop()
	player.OMX.DisconnectPlayer()
}
