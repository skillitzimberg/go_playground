// Hands-on exercise #7
// write a program that:
// 	- launches 10 goroutines:
// 		* each goroutine adds 10 numbers to a channel
//  - pull the numbers off the channel and print them

package main

import (
	"fmt"
	"sync"
)

func exerciseSeven() {
	var wait sync.WaitGroup
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			wait.Add(1)
			go func(m int) {
				for j := 0; j < 10; j++ {
					c <- j*10 + m
				}
				wait.Done()
			}(i)
		}
		wait.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("About to exit main routine.")
}
