package main

import "fmt"

func main() {
	fmt.Println(divisibleByFourSum(sum, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16))
}

func sum(item ...int) int {
	total := 0

	for _, value := range item {
		total += value
	}

	return total
}

func divisibleByFourSum(f func(v ...int) int, x ...int) int {
	var y []int
	for _, a := range x {
		if a%4 == 0 {
			y = append(y, a)
		}
	}

	return f(y...)
}
