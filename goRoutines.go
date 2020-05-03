package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func goRoutines() {
	fmt.Println("BEGIN")
	fmt.Println("Beginning CPU", runtime.NumCPU())
	fmt.Println("Beginning Routines", runtime.NumGoroutine())

	wg.Add(2)
	go func() {
		fmt.Println("Begin Thing One")
		for i := 0; i < 5; i++ {
			fmt.Printf("Thing One %d\n", i)
		}
		wg.Done()
	}()

	go func() {
		fmt.Println("Begin Thing Two")
		for i := 7; i < 30; i++ {
			fmt.Printf("Thing Two %d\n", i)
		}
		wg.Done()
	}()

	func() {
		fmt.Println("Begin Thing Three")
	}()

	fmt.Println("Pre-wait CPU", runtime.NumCPU())
	fmt.Println("Pre-wait Routines", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Ending CPU", runtime.NumCPU())
	fmt.Println("Ending Routines", runtime.NumGoroutine())
	fmt.Println("END")
}
