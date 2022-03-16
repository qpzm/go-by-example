package main

import "fmt"

func main() {
    // if the buffer size is 1, an error is raised.
    // fatal error: all goroutines are asleep - deadlock!
    messages := make(chan string, 2)
    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
