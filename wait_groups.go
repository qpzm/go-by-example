package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)

        // Create a new variable
        i := i

        go func() {
            defer wg.Done() // decrement the wait group counter by 1
            worker(i)
        }()
    }

    wg.Wait()
}
