package main

import "fmt"

func main() {
	i := 10
	fmt.Printf("%d\t%b\t%#x", i, i, i)

	j := i << 1
	fmt.Printf("\n%d\t%b\t%#x", j, j, j)
}
