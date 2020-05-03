package main

func routine(routineID int) {
	go func() {
		mux.Lock()
		value := counter
		// runtime.Gosched()
		value++
		counter = value
		mux.Unlock()

		Wait.Done()
	}()
}
