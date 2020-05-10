package main

import "fmt"

func directionalChannels() {
	bdc := make(chan int)
	var num int

	go sendChannel(bdc)
	go receiveChannel(bdc, &num)

	fmt.Println("num:", num)

	fmt.Printf("------\n")
	fmt.Printf("bdc\t%T\n", bdc)
}

func sendChannel(cs chan<- int) {
	cs <- 42
}

func receiveChannel(cr <-chan int, num *int) {
	*num = <-cr
}
