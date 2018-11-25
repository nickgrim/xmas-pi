package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

const delay = time.Second / 2

// GPIOs for the LEDs
var star rpio.Pin = 2
var leds = []rpio.Pin{4, 15, 13, 21, 25, 8, 5, 10, 16, 17, 27, 26, 24, 9, 12, 6, 20, 19, 14, 18, 11, 7, 23, 22} // …in this order

func main() {
	var (
		done bool
	)
	// Set up GPIO, and set all of our pins to be an output
	if err := rpio.Open(); err != nil {
		log.Fatalln(err)
	}
	star.Output()
	for _, p := range leds {
		p.Output()
	}
	defer rpio.Close()
	defer turnOffLEDs()
	// Trap SIGINT and SIGTERM, and set done
	sigs := make(chan os.Signal, 1)
	go func() {
		_ = <-sigs
		done = true
	}()
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// Turn on the star, and blink random LEDs until we're done
	star.High()
	for !done {
		randomlySetLEDs()
		time.Sleep(delay)
	}
}

// randomlySetLEDs iterates through the LEDs, setting each to a random high/low state.
func randomlySetLEDs() {
	rnd := rand.Uint32()
	for _, p := range leds {
		if rnd&1 == 1 {
			p.Toggle()
		}
		rnd = rnd >> 1
	}
}

// turnOffLEDs turns off all of the LEDs by setting the GPIO low.
func turnOffLEDs() {
	star.Low()
	for _, p := range leds {
		p.Low()
	}
}
