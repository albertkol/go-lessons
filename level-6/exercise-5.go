package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (s circle) area() float64 {
	return math.Pi * s.radius * s.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	square := square{10,20}
	circle := circle{15}

	info(square)
	info(circle)
}
