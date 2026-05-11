package main

import "fmt"

func greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

func main() {
    fmt.Println("Go Project test")
    fmt.Println(" ", greet("mide"))
}
