package main

import "fmt"

func main() {
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    // 1. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
    // 2. it’s possible to close a non-empty channel but still have the remaining values be received.
    for elem := range queue {
        fmt.Println(elem)
    }
}
