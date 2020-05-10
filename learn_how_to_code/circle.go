package main

import (
	"math"
)

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (c *circle) getAddress() **circle {
	return &c
}
