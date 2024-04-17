package main

import (
    "fmt"
    "time"
)

func main() {
    // Measure the time it takes to execute the function
    start := time.Now()

    // Call the function you want to measure
    yourFunction()

    // Calculate the elapsed time
    elapsed := time.Since(start)

    // Format the elapsed time to milliseconds with high precision
    elapsedMilliseconds := elapsed.Seconds() * 1000

    // Convert the elapsed time to a string
    elapsedString := fmt.Sprintf("%.7f", elapsedMilliseconds)

    fmt.Println("Elapsed Time:", elapsedString, "milliseconds")
}

// The function you want to measure
func yourFunction() {
    // Simulate some work
    time.Sleep(1 * time.Second)
}
