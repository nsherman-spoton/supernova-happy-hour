# supernova-happy-hour
A space to create and enjoy

## GoBots
Gobot is a framework for robotics, physical computing, and the Internet of Things (IoT), written in the Go programming language.

Gobot provides drivers and adapters for controlling a wide variety of physical devices from low-level Arduino and Raspberry Pi, as well as drones, toys, and other complete devices that themselves have APIs.

https://gobot.io/

### Interesting platforms
https://gobot.io/documentation/platforms/

- NATS
- Arduino
- LEAP Motion


## Setup Arduino
- In Arduino IDE go to library and install Firmata package
- From the Firmata package (in the library window) use the '...' menu to open the example StandardFirmata
- Upload the sketch to the arduino

## Setup Code Runner / Server
- connect configured arduino to an Arch: amd64 machine
- pull repo and get the configuration for the connection to the serial device and nats cluster
- run go code by entering the directory and running `NATS_URL={val} NATS_USER={val} NATS_PASS={val} ARDUINO_PATH={val} go run .` 

## Setup NATS Cli
- install the nats cli `go install github.com/nats-io/natscli/nats@latest`
- add the context `nats context add testing --server={hostname} --user={user} --password={pass}`
- select the context `nats context select testing`
- publish messages `nats pub "lights:on" red` or `nats pub "lights:off" green`
