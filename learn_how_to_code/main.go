package main

import (
	"fmt"

	"github.com/skillitzimberg/go_playground/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	pete := canine{
		"Pete", dog.Years(49),
	}

	fmt.Println(pete)
}
