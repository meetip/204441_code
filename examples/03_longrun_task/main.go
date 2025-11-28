package main

import (
	"fmt"
	"time"
)

func longTask() {
	// Simulate a long-running task
	time.Sleep(1 * time.Minute)
	fmt.Println("Long task finished")
}

func main() {
	go longTask()

	fmt.Println("Started long task (will take ~1m)")

	// Main waits 30 seconds and then exits without waiting for the long task
	time.Sleep(30 * time.Second)
	fmt.Println("Main finished (did not wait for long task)")
}
