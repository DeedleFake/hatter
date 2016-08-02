package main

import (
	"time"
)

const (
	Base = "/etc/hatter"
)

func main() {
	tick := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-tick.C:
		}
	}

	panic("This should never run.")
}
