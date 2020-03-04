package main

import "fmt"

func main() {
	v := []int{1,2,3,4,5,6,7,8,9}

	fmt.Println(sum(v...))
	fmt.Println(sum2(v))
}

func sum(v ...int) int {
	sum := 0
	for _, b := range v {
		sum += b
	}

	return sum
}

func sum2(v []int) int {
	sum := 0
	for _, b := range v {
		sum += b
	}

	return sum
}
