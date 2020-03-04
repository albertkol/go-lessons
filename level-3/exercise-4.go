package main

import "fmt"

func main() {
	i:=1993
	for {
		fmt.Println(i)

		i++
		if i == 2021 {
			break
		}
	}
}
