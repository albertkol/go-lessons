package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go increment("First:")
	go increment("Second:")
	wg.Wait()
	fmt.Println(counter)
}

func increment(s string) {
	for i := 0; i < 5; i++ {
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter))
	}
	wg.Done()
}
