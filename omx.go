package main

import (
//	"bufio"
	"fmt"
	"time"
//	"os"
//	"strconv"
//	"strings"

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
//		fmt.Println("Player active")
		status, err := player.OMX.PlaybackStatus()
		if err != nil {
			fmt.Println("Stop error: ", err)
		}
		if status == "playing" {
//			fmt.Println("Still playing")
			player.OMX.Pause()
			time.Sleep(time.Duration(250) * time.Millisecond)
		}
	}
}

func (player *OMXPlayer)Start() {
	if player.OMX.HasPlayer() == false {
                fmt.Println("Start playing")
		player.OMX.StartPlayer(player.Media.NextTrack())
		time.Sleep(time.Duration(250) * time.Millisecond)
        } else {
//		fmt.Println("Player startet")
		status, err := player.OMX.PlaybackStatus()
                if err != nil {
                        fmt.Println("PlaybackStatus error: ", err)
                }
                if status != "playing" {
                        fmt.Println("Restart playing")
			player.OMX.Play()
			time.Sleep(time.Duration(250) * time.Millisecond)		
		} else {
//			fmt.Println("Allready playing")
		}
	}

}

func (player *OMXPlayer)Close() {
	player.OMX.Stop()
	player.OMX.DisconnectPlayer()
}
