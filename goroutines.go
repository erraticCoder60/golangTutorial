package main

import (
	"fmt"
	"time"
)

// thread

func firstGoroutine() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration in firstGoroutine: ", i+1)
		time.Sleep(1 * time.Millisecond)
	}
}

func secondGoroutine() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration in secondGoroutine: ", i+1)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Start a goroutine to execute firstGoroutine concurrently
	go firstGoroutine()

	// Start a goroutine to execute secondGoroutine concurrently
	go secondGoroutine()

	//anonymous goroutines
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Iteration in thirdGoroutine: ", i+1)
			time.Sleep(1 * time.Millisecond)
		}
	}()

	// Wait for a short duration to allow the goroutine to execute
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Main function")
}
