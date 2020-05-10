package main

import "fmt"

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (p person) intro() {
	fmt.Printf("My name is %s %s and I am %d years old.\n", p.First, p.Last, p.Age)
}

func (p *person) speak() {
	fmt.Printf("My name is %s %s.\n", p.First, p.Last)
}
