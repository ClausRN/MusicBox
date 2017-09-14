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
)

func main() {

	var name string
	name="/usr/bin/omxplayer"

	cmd := exec.Command(name, "-version")
        fmt.Println("Staring player")
	err := cmd.Run()

	fmt.Println("process launched")

	if err!=nil {
		var buffer bytes.Buffer
		buffer.WriteString("Could not start omxplayer: ")
		buffer.WriteString(err.Error())
//		errortext = "Could not start omxplayer: " + err
		fmt.Println("Could not start omxplayer")
		panic(buffer.String())
	}

        fmt.Println("Player startet")

	LedOn := false

	h := NewUSProbe(3, 2, 15)		// Echo pin, trigger pin, timeout
	Play := NewOMXPlayer()

	pir := NewPIR(21)

        if err := rpio.Open(); err != nil {
                panic(err.Error())
        }
        LEDPin.Output()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	defer signal.Stop(quit)

	for true {
                //fmt.Println("Probing")
		h.Probe()
//		fmt.Printf("Time %v: Distance: %2.3f, Min: %2.3f, Max: %2.3ff, Raw: %2.3f  \n", h.TriggeredUntil, h.Distance, h.DistanceLow, h.DistanceHigh, h.DistanceRaw)
//		if (h.TriggeredUntil.After(time.Now())){
                if (pir.IsActive()) {
			//fmt.Println("Playing")
			Play.Start()
		} else {
			Play.Stop()	
		}
		//fmt.Println("Check signals")
		select {
		case <-time.After(200 * time.Millisecond):
			if LedOn {
				LEDPin.High()
				LedOn = false
			} else {
				LEDPin.Low()
				LedOn = true
			}
			//fmt.Println("Pause over")
		case <-quit:
			fmt.Println("Catch ctrl-c, exiting!")
			Play.Close()
			LEDPin.Low()
			return
		}
		//time.Sleep(time.Duration(250) * time.Millisecond)
	}

}
