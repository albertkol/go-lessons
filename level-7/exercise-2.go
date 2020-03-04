package main

import "fmt"

type person struct {
	first string
	second string
}

func changeMe(p *person) {
	p.first = "Not Albert"
}

func main() {
	person := person{
		first: "Albert",
		second: "Kolozsvari",
	}

	fmt.Printf("%T\n%v\n", person, person)
	changeMe(&person)
	fmt.Printf("%T\n%v\n", person, person)
}
