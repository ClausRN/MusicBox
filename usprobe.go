package main

import (
	"time"
	"github.com/stianeikeland/go-rpio"
)

const HardStop = 1000000
const USSensitivity = 5

type USProbe struct {
	EchoPin 	rpio.Pin
	PingPin 	rpio.Pin
	TriggeredUntil 	time.Time
	Timeout		int
        Distance     	float32
	DistanceLow	float32
	DistanceHigh	float32
	DistanceRaw	float32
}

func NewUSProbe(echo    int, // Echo pin
	        ping    int, // Trigger pin
		Timeout int) (result USProbe) {
	if err := rpio.Open(); err != nil {
		panic(err.Error())
	}
	result.EchoPin = rpio.Pin(echo)
	result.PingPin = rpio.Pin(ping)
	result.TriggeredUntil = time.Now().Add(-(time.Duration(1)  * time.Microsecond))
	var first  float32
        var second float32
        var third  float32
 	result.Timeout = Timeout
	first = result.MeasureDistance()
        second = result.MeasureDistance()
        third = result.MeasureDistance()
	
	result.Distance = (first + second + third)/3
	result.DistanceLow = result.Distance - (result.Distance * float32(USSensitivity)) / float32(100)
        result.DistanceHigh = result.Distance + (result.Distance * float32(USSensitivity)) / float32(100)
	return
}

func (hcsr *USProbe) MeasureDistance() float32 {

	hcsr.EchoPin.Output()
	hcsr.PingPin.Output()

	hcsr.EchoPin.Low()
	hcsr.PingPin.Low()

	hcsr.EchoPin.Input()

	strobeZero := 0
	strobeOne := 0

	// strobe
	hcsr.delayUs(200)
	hcsr.PingPin.High()
	hcsr.delayUs(10)
	hcsr.PingPin.Low()

	// wait until strobe back

	for strobeZero = 0; strobeZero < HardStop && hcsr.EchoPin.Read() != rpio.High; strobeZero++ {
	}
	start := time.Now()
	for strobeOne = 0; strobeOne < HardStop && hcsr.EchoPin.Read() != rpio.Low; strobeOne++ {
		hcsr.delayUs(1)
	}
	end := time.Now()
	hcsr.DistanceRaw = float32(end.UnixNano()-start.UnixNano()) / (58.0 * 1000)
	return hcsr.DistanceRaw
}

func (hcsr *USProbe) delayUs(ms int) {
	time.Sleep(time.Duration(ms) * time.Microsecond)
}

func (hcsr *USProbe) Probe(){
        var first  float32
	first = hcsr.MeasureDistance()
	if (hcsr.DistanceRaw > hcsr.DistanceHigh) || (hcsr.DistanceRaw < hcsr.DistanceLow){
		hcsr.TriggeredUntil = time.Now().Add(time.Duration(hcsr.Timeout) * time.Second)
        	var second float32
	        var third  float32

	        second = hcsr.MeasureDistance()
        	third = hcsr.MeasureDistance()

	        hcsr.Distance = (first + second + third)/3
        	hcsr.DistanceLow = hcsr.Distance - (hcsr.Distance * float32(USSensitivity)) / float32(100)
	        hcsr.DistanceHigh = hcsr.Distance + (hcsr.Distance * float32(USSensitivity)) / float32(100)
	}
}

