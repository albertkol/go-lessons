package main

import "fmt"

const (
	x = 2020 + iota
	y = 2020 + iota
	z = 2020 + iota
	t = 2020 + iota
)

const (
	a = 2020 + iota
	b = 2020 + iota
	c = 2020 + iota
	d = 2020 + iota
)

func main() {
	fmt.Println(x, y, z, t)
	fmt.Println(a, b, c, d)
}
