package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	fmt.Println(x)

	for _, a2 := range x {
		for _, b2 := range a2 {
			fmt.Println(b2)
		}
	}
}
