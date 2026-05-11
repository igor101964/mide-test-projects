package main
import "fmt"
func main() {
    fmt.Println("Hello from Go!")
    for i := 1; i <= 3; i++ {
        fmt.Printf("  step %d\n", i)
    }
}
