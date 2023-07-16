package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

// Program that prints the current time / exact time using the NTP library
func main() {
	println("Start")
	CurrentTime()
	println("End")
	os.Exit(0)
}
func CurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Current time: %v\n", time)
}
