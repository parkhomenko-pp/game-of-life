package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.High()
		println("high")
		time.Sleep(time.Second)

		led.Low()
		println("low")
		time.Sleep(time.Second)
	}
}
