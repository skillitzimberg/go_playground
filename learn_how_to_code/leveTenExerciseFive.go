// Show the comma ok idiom.

package main

import (
	"fmt"
)

func levelTemExerciseFive() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)
	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
