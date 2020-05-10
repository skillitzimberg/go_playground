package main

import (
	"fmt"
)

func selectFromChannels() {
	q := make(chan bool) // Make a bidirectional channel q
	c := addToChannel(q)

	receiveFromChannels(c, q)

	fmt.Println("about to exit")
}

func addToChannel(q chan<- bool) <-chan int {
	c := make(chan int) // Make a bidirectional channel c
	go func() {         // Launch a Go routine
		for i := 0; i < 100; i++ {
			c <- i // Add values to the channel c
		}
		close(c)
		_, ok := <-c // Check that channel c is closed.
		q <- ok      // Add a value to the channel q
	}()
	return c // Return c as a receive channel
}

func receiveFromChannels(c <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case g := <-q:
			fmt.Println(g)
			return
		}
	}
}
