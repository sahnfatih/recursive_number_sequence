package main

import (
    "fmt"
    "os"
    "strconv"
    "recursive_number_sequence/internal/sequence"
)

func main() {
    // Check if argument is provided
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run cmd/main.go <number>")
        fmt.Println("Example: go run cmd/main.go 9")
        os.Exit(1)
    }

    // Parse input number
    num, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Printf("Error: Invalid number format - %v\n", err)
        os.Exit(1)
    }

    // Generate sequence
    result := sequence.GenerateSequence(num)

    // Print results
    fmt.Printf("Input: %d\n", num)
    fmt.Printf("Output: ")
    for _, v := range result {
        fmt.Printf("%d\n", v)
    }
}