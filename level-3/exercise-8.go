package main

import "fmt"

func main() {
	for i:=10; i<=100; i++ {
		switch {
		case i%4 == 0:
			fmt.Println(i, " is divisible with 4")
		case i%4 == 1:
			fmt.Println(i, " remainder is 1")
		case i%4 == 2:
			fmt.Println(i, " remainder is 2")
		case i%4 == 3:
			fmt.Println(i, " remainder is 3")
		}
	}
}
