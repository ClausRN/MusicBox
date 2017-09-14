package main

import (
	"fmt"
        "github.com/stianeikeland/go-rpio"
)

type PIR struct {
	PIRPin rpio.Pin
}

func NewPIR(PIRPin int) (result PIR) {
	if err := rpio.Open(); err != nil {
		panic(err.Error())
	}
	
	result.PIRPin = rpio.Pin(PIRPin)
	result.PIRPin.Input()
	fmt.Println("Open for input: ", PIRPin)
	return
}

func (sensor *PIR)IsActive() bool {
	if sensor.PIRPin.Read() != rpio.Low {
//		fmt.Println("High")
		return true
	} else {
//		fmt.Println("Low")
		return false
	}

}
