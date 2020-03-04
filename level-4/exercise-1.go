package main

import "fmt"

func main() {
	var a [5]int
	a[0]=1
	a[1]=2
	a[2]=3
	a[3]=4
	a[4]=5

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", a)
}
