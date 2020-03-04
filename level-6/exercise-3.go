package main

import "fmt"

func main() {
	defer first()
	second()
}

func first() {
	fmt.Println("First!")
}

func second() {
	fmt.Println("Second!")
}
