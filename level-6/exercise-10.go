package main

import "fmt"

func main() {
	a := increment()
	b := increment()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(a())
}

func increment() func() int {
	x := 1

	return func() int {
		x += x

		return x
	}
}
