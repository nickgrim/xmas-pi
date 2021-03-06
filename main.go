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
		ticksChan  = make(chan time.Time)
		toggleChan = make(chan bool)
		quitChan   = make(chan struct{})
	)

	// Trap SIGHUP, SIGINT, SIGTERM, SIGUSR2
	sigs := make(chan os.Signal, 1)
	go func() {
		for {
			switch <-sigs {
			case syscall.SIGUSR2:
				toggleChan <- true
			default:
				close(quitChan)
				return
			}
		}
	}()
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)

	// Initialise GPIO
	if err := rpio.Open(); err != nil {
		log.Fatalln(err)
	}
	defer rpio.Close()
	// Set all of our pins to be an output
	star.Output()
	for _, p := range leds {
		p.Output()
	}
	defer turnOffLEDs()

	// Start generating ticks
	go generateTicks(ticksChan, toggleChan)

	// Turn on the star, and blink random LEDs until we're done
	for {
		select {
		case <-quitChan:
			return
		case <-ticksChan:
			randomlySetLEDs()
		}
	}
}

// randomlySetLEDs iterates through the LEDs, setting each to a random high/low state.
func randomlySetLEDs() {
	star.High()
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

// generateTicks uses a timer to generate regular ticks into tickerChan. It can be toggled via toggleChan.
func generateTicks(tickerChan chan<- time.Time, toggleChan <-chan bool) {
	var (
		ticker    = time.NewTicker(delay)
		sendTicks = true
	)

	for {
		select {
		case t := <-ticker.C:
			if sendTicks {
				tickerChan <- t
			}
		case <-toggleChan:
			sendTicks = !sendTicks
			turnOffLEDs()
		}

	}
}
