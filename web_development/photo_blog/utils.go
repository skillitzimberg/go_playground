package main

import (
	"fmt"
	"log"
)

func check(err error, from string) {
	if err != nil {
		s := fmt.Sprintf("Error from %s: %s", from, err.Error())
		log.Println(s)
	}
}
