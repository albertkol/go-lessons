package main

import "fmt"

func main() {
	x := struct {
		first string
		last  string
	}{
		first: "Joe",
		last:  "Doe",
	}

	fmt.Println(x)
}
