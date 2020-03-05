package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		fmt.Println("First routine")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second routine")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("All done!")
}
