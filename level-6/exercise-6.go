package main

import "fmt"

func main() {
	fmt.Println(func(a int, b int) int {
		return a + b
	}(1, 3))
}
