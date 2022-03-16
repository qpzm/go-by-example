package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var ops uint64

    // A WaitGroup will help us wait for all goroutines to finish their work.
    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {
                // If we use ops++, the result <= 50000
                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait() // wait until the counter is zero

    // To read atomically, use atomic.LoadUint64.
    fmt.Println("ops:", ops)
}
