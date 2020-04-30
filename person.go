package main

import "fmt"

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (p person) intro() {
	fmt.Printf("My name is %s %s and I am %d years old.", p.First, p.Last, p.Age)
}
