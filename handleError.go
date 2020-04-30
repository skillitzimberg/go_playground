package main

import "fmt"

func fakeHandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
