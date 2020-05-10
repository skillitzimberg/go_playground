package main

import (
	"fmt"
	"runtime"
)

func printSystemVariables() {
	fmt.Println("GO ARCH:", runtime.GOARCH)
	fmt.Println("GO OS:", runtime.GOOS)
}
