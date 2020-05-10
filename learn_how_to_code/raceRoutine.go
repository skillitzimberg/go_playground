package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var Wait sync.WaitGroup
var mux sync.Mutex
var counter int64

func routine(routineID int) {
	fmt.Println("Goroutines begin:", runtime.NumGoroutine())

	const routineMax int = 100

	Wait.Add(routineMax)

	for i := 0; i < routineMax; i++ {

		go func() {
			// mux.Lock()
			atomic.AddInt64(&counter, 1)
			fmt.Println("COUNTER AFTER ADD:", atomic.LoadInt64(&counter))
			// runtime.Gosched()
			// mux.Unlock()

			Wait.Done()
		}()

		fmt.Printf("Goroutines running on run #%d: %d\n", i+1, runtime.NumGoroutine())
	}

	Wait.Wait()
	fmt.Println("Goroutines end:", runtime.NumGoroutine())
	fmt.Println("Final Counter:", counter)

}
