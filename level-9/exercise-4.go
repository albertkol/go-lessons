package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mx sync.Mutex

func main() {
	wg.Add(2)
	go increment("First:")
	go increment("Second:")
	wg.Wait()
	fmt.Println(counter)
}

func increment(s string) {
	for i:=0; i<5; i++ {
		mx.Lock()
		counter++
		fmt.Println(s,i, "Counter:", counter)
		mx.Unlock()
	}
	wg.Done()
}
