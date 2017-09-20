package main

import (
	"time"
        "github.com/stianeikeland/go-rpio"
)

type PIR struct {
	PIRPin 		rpio.Pin
        TriggeredUntil  time.Time
        Timeout         int
}

func NewPIR(PIRPin int, Timeout int) (result PIR) {
	if err := rpio.Open(); err != nil {
		panic(err.Error())
	}
	result.PIRPin = rpio.Pin(PIRPin)
	result.PIRPin.Input()
        result.TriggeredUntil = time.Now().Add(-(time.Duration(1)  * time.Microsecond))
        result.Timeout = Timeout
	return
}

func (sensor *PIR)IsActive() bool {
	if sensor.PIRPin.Read() != rpio.Low {
		sensor.TriggeredUntil = time.Now().Add(time.Duration(sensor.Timeout) * time.Second)
	}
	if (sensor.TriggeredUntil.After(time.Now())){
		return true
	} else {
		return false
	}

}
