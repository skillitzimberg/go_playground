package main

import "fmt"

func levelTenExcerciseSeven() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func(k int) {
			for j := 0; j < 10; j++ {
				c <- j + 10*k
			}
		}(i)
	}

	for m := 0; m < 100; m++ {
		fmt.Println(m, <-c)
	}
}
