package main

import "fmt"

// Go Routine main calls basicChannel()
func basicChannel() {
	// Go Routine main makes channel c of type int
	c := make(chan int)

	// Go Routine main launches Go Routine two
	go func() {
		// Go Routine two put the int 42 onto channel c
		c <- 42
		// Channel c blocks writing to channel c until it can be read by Go Routine main.
	}()

	// Go Routine main is blocked from reading channel c until Go Routine two is ready to write to it.

	// When Go Routines main and two are sync ready to read & write respectively the actions occur simultaneously. Like a baton pass in a relay race.
	fmt.Println(<-c)
}
