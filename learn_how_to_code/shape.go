package main

import "fmt"

type shape interface {
	area() float64
}

func shapeInfo(s shape) {
	fmt.Println("Area:", s.area())
}

func shapeAddress(s shape) {
	fmt.Println("Address of the shape itself:", &s)
}

func shapeValue(s shape) {
	fmt.Println("Address of the shape itself:", *&s)
}
