package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: mapture <command>")
        return
    }

    command := os.Args[1]

    switch command {
    case "list":
        fmt.Println("Listing keymaps... (soon!)")
    case "search":
        fmt.Println("Searching keymaps... (soon!)")
    default:
        fmt.Printf("Unknown command: %s\n", command)
    }
}
