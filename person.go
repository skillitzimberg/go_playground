package main

import "fmt"

type person struct {
	surname   string
	firstname string
	age       int
}

func (p person) intro() {
	fmt.Printf("My name is %s %s and I am %d years old.", p.firstname, p.surname, p.age)
}
