package main

import (
	"fmt"
	"os"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/firmata"
	"gobot.io/x/gobot/v2/platforms/nats"
)

func main() {
	natsAdaptor := nats.NewAdaptorWithAuth(os.Getenv("NATS_URL"), 1, os.Getenv("NATS_USER"), os.Getenv("NATS_PASS"))
	firmataAdaptor := firmata.NewAdaptor(os.Getenv("ARDUINO_PATH"))
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
