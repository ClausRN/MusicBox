package main

import (
	"fmt"
	"bytes"
	"time"
	"os"
	"os/signal"
	"os/exec"
	"github.com/stianeikeland/go-rpio"
)

var (
        LEDPin rpio.Pin = rpio.Pin(26)
	PirPin int = 21
	Timeout int = 10
)

func main() {

	var name string
	name="/usr/bin/omxplayer"

	cmd := exec.Command(name, "-version")
        fmt.Println("Starting player")
	err := cmd.Run()
	fmt.Println("  Process launched")
	if err!=nil {
		var buffer bytes.Buffer
		buffer.WriteString("Could not start omxplayer: ")
		buffer.WriteString(err.Error())
		fmt.Println("Could not start omxplayer")
		panic(buffer.String())
	}
        fmt.Println("  Player startet")

	LedOn := false
	Play := NewOMXPlayer()
	pir := NewPIR(PirPin, Timeout)
        if err := rpio.Open(); err != nil {
                panic(err.Error())
        }
        LEDPin.Output()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	defer signal.Stop(quit)

	for true {
                if (pir.IsActive()) {
			Play.Start()
		} else {
			Play.Stop()	
		}
		select {
		case <-time.After(200 * time.Millisecond):
			if LedOn {
				LEDPin.High()
				LedOn = false
			} else {
				LEDPin.Low()
				LedOn = true
			}
		case <-quit:
			fmt.Println("Catch ctrl-c, exiting!")
			Play.Close()
			LEDPin.Low()
			return
		}
	}

}
