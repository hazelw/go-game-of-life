package main

import (
    "fmt"
    "time"
    "math/rand"
)

// "When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last"
func get_random_number(min, max int) int {
    // We need to seed the random number generator with the current datetime
    // as without it Go uses a shared Source(???) to generate the same
    // deterministic sequence of values every time this program is run.

    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min + 1) + min
}

// "A function can return any number of results"
func get_three_things() (string, string, string) {
    return "thing 1", "thing 2", "thing 3"
}

func main() {
    fmt.Println("hello, world\n")
    fmt.Println("The time is currently ", time.Now())
    fmt.Println("Here is a random number between 3 and 10: ", get_random_number(3, 10))

    var first_thing, second_thing, third_thing string

    first_thing, second_thing, third_thing = get_three_things()
    fmt.Println("Here is the first thing: ", first_thing)
    fmt.Println("Here is the second thing: ", second_thing)
    fmt.Println("Here is the third thing: ", third_thing)
}
