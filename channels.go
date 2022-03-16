package main

import "fmt"

func main() {
    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages // block until both the sender and receiver are ready
    fmt.Println(msg)
}
