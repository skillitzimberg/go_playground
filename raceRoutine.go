package main

import (
	"fmt"
	"sync/atomic"
)

func routine(routineID int) {
	go func() {
		// mux.Lock()
		atomic.AddInt64(&counter, 1)
		fmt.Println("COUNTER AFTER ADD:", atomic.LoadInt64(&counter))
		// runtime.Gosched()
		// mux.Unlock()

		Wait.Done()
	}()
}
