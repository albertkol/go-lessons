package main

import "fmt"

func main() {
	a := first()

	a()
}

func first() func() {
	return func() {
		fmt.Println("Hello!")
	}
}
