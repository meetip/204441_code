package main

import (
	"fmt"
	"sync"
	"time"
)

func longTask(wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate a long-running task
	time.Sleep(1 * time.Minute)
	fmt.Println("Long task finished")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go longTask(&wg)

	fmt.Println("Started long task (will take ~1m)")

	// Main can do other work here; then wait for the long task to finish
	time.Sleep(30 * time.Second)
	fmt.Println("Main waited 30s, now waiting for long task to finish...")

	wg.Wait()
	fmt.Println("All done")
}
