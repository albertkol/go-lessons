package main

import "fmt"

var x int

func main() {
	fmt.Println(&x)

	fmt.Printf("%T\n", &x)
	fmt.Printf("%T\n", *&x)
}
