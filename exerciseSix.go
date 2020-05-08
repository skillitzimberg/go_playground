// Hands-on exercise #6
// write a program that:
// 	- puts 100 numbers to a channel
//  - pull the numbers off the channel and print them

package main

import "fmt"

func exerciseSix() {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	fmt.Println("About to exit main.")
}
