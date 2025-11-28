package main

import (
    "fmt"
    "sync"
)

// Fixed race condition: protect the shared counter with a mutex.
// Alternative: use atomic.AddInt32/atomic.AddInt64 for lock-free increments.

func main() {
    var wg sync.WaitGroup 
    count := 0

    wg.Add(1000)
    for i := 0; i < 1000; i++ {
        go func() {
            defer wg.Done() 
            count++ 
        }()
    }

    wg.Wait()
    fmt.Printf("Final count: %d\n", count)
}
