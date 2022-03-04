package main

import (
    "fmt"
    "os"
)

func main() {
    defer fmt.Println("1") // 1 is not printed
    os.Exit(3)
}
