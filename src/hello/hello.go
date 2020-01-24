package main

import (
    "fmt"
    "time"
    "math/rand"
)

func get_random_number(min int, max int) int {
    // We need to seed the random number generator with the current datetime
    // as without it Go uses a shared Source(???) to generate the same
    // deterministic sequence of values every time this program is run.

    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min + 1) + min
}

func main() {
    fmt.Println("hello, world\n")
    fmt.Println("The time is currently ", time.Now())
    fmt.Println("Here is a random number between 3 and 10: ", get_random_number(3, 10))
}
