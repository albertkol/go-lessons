package main

import "fmt"

type person struct {
	first string
	last string
}

func (p *person) speak() {
	fmt.Println("Hello, I am", p.first, p.last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{"Albert", "Kolozsvari"}

	saySomething(&p)
}
