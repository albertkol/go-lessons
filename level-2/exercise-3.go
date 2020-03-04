package main

import "fmt"

const typed = 10;
const untyped int = 11;

const (
	hello = "hello"
	world string = "world"
)

func main() {
	fmt.Println(typed, untyped, hello, world)
}
