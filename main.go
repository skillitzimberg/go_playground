package main

import (
	"fmt"
	"sync"
)

var Wait sync.WaitGroup
var Counter int = 0

func main() {

	for routine := 1; routine <= 2; routine++ {

		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {
	fmt.Println("routine ID:", id)
	for count := 0; count < 2; count++ {

		value := Counter
		value++
		Counter = value
		fmt.Printf("routine ID: %d; routine counter: %d", id, Counter)
	}

	Wait.Done()
}
