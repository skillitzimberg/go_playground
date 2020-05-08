// Hands-on exercise #1
// 	get this code working:
// 		- using func literal, aka, anonymous self-executing func
// 		- a buffered channel

package main

import (
	"fmt"
)

func levelTenExerciseOneA() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func levelTenExerciseOneB() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
