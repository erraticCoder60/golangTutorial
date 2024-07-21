package main

import "fmt"

// This function accepts a variable number of integers as arguments
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// Calling the sum function with different numbers of arguments
	fmt.Println(sum())              // Outputs: 0
	fmt.Println(sum(1, 2))          // Outputs: 3
	fmt.Println(sum(1, 2, 3))       // Outputs: 6
	fmt.Println(sum(1, 2, 3, 4, 5)) // Outputs: 15
}
