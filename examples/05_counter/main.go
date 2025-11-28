package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func(val int) {
			fmt.Printf("%d ", val)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
