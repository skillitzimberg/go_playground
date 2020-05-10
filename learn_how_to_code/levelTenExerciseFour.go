// Pull the values off the channel using a select statement.

package main

import (
	"fmt"
)

func levelTenExerciseFour() {
	q := make(chan int)
	c := genFour(q)

	receiveFour(c, q)

	fmt.Println("about to exit")
}

func genFour(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()

	return c
}

func receiveFour(cr, q <-chan int) {
	for {
		select {
		case v := <-cr:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
