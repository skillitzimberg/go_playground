package main

import (
	"fmt"
	"runtime"
	"sync"
)

var Wait sync.WaitGroup
var mux sync.Mutex
var counter int = 0

func main() {
	fmt.Println("Goroutines begin:", runtime.NumGoroutine())

	const routineMax int = 100

	Wait.Add(routineMax)

	for i := 0; i < routineMax; i++ {

		go routine(i)

		fmt.Printf("Goroutines running on run #%d: %d\n", i+1, runtime.NumGoroutine())
	}

	Wait.Wait()
	fmt.Println("Goroutines end:", runtime.NumGoroutine())
	fmt.Println("Final Counter:", counter)
}
