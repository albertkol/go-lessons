package main

import "fmt"

func main() {
	for i:=10; i<=100; i++ {
		if i%4 == 0 {
			fmt.Println(i, " is divisible with 4")
		} else if i%4 == 1 {
			fmt.Println(i, " remainder is 1")
		} else if i%4 == 2 {
			fmt.Println(i, " remainder is 2")
		} else if i%4 == 3 {
			fmt.Println(i, " remainder is 3")
		}
	}
}
