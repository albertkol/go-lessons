package main

import "fmt"

func main() {
	g := 3.0 == 3
	h := 4 <= 5
	i := 10 >= 5.5
	j := 4 != 4.0
	k := 3 < 4
	l := 5 > 9

	fmt.Printf("%t\t%t\t%t\t%t\t%t\t%t", g, h, i, j, k, l)
}
