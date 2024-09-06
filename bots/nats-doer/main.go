package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot/platforms/nats"
)

func main() {
	natsAdaptor := nats.NewAdaptor("localhost:4222", 1)
	firmataAdaptor := firmata.NewAdaptor()
	goodLED := gpio.NewLedDriver(firmataAdaptor, "12")
	evilLED := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		natsAdaptor.On("lights:on", func(msg nats.Message) {
			color := string(msg.Data)

			var led *gpio.LedDriver
			if color == "red" {
				led = evilLED
			} else {
				led = goodLED
			}

			if led.State() {
				return
			}

			if err := led.On(); err != nil {
				fmt.Printf("Error turning on LED: %v", err)
			}
		})

		natsAdaptor.On("lights:off", func(msg nats.Message) {
			color := string(msg.Data)

			var led *gpio.LedDriver
			if color == "red" {
				led = evilLED
			} else {
				led = goodLED
			}

			if !led.State() {
				return
			}

			if err := led.Off(); err != nil {
				fmt.Printf("Error turning off LED: %v", err)
			}
		})
	}

	robot := gobot.NewRobot("nats-doer",
		[]gobot.Connection{natsAdaptor, firmataAdaptor},
		[]gobot.Device{goodLED, evilLED},
		work,
	)

	robot.Start()
}
