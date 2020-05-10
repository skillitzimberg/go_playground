// Pull the values off the channel using a for range loop.

package main

import (
	"fmt"
)

func levelTenExerciseThree() {
	c := generate()
	readChannel(c)

	fmt.Println("about to exit")
}

func generate() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func readChannel(cr <-chan int) {
	for v := range cr {
		fmt.Println(v)
	}
}
