package main

import "fmt"

// Go Routine Main launches
func channelBuffer() {
	// Go Routine Main creates channel c
	// Channel c is a buffer. It is capable of holding 1 value of type int.
	c := make(chan int, 1)
	// Although channels are blocking c's ability to buffer (hold that one value) means that GO Routine Main can write to the channel without waiting for another Go routine to be ready to read from the channel.
	c <- 42
	// And Go Routine Main can then take the value off of channel c.
	fmt.Println(<-c)

	// Buffer channels are a way to use channels without multiple Go routines running.
}
