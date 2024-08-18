package main

import "fmt"

func main() {
    // Create a map with keys as integers and values as empty structs
    seen := make(map[int]struct{})

    // Example number to store in the map
    num := 1337

    // Store the number in the map with an empty struct as the value
    seen[num] = struct{}{}

    // Check if the number is in the map
    if _, exists := seen[num]; exists {
        fmt.Println("Number", num, "exists in the map!")
    } else {
        fmt.Println("Number", num, "does not exist in the map.")
    }
}
