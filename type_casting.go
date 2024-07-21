package main

import (
	"fmt"
)

func main() {
	// Example 1: Convert int to float64
	var a int = 10
	var b float64 = float64(a)
	fmt.Printf("a: %d, b: %f\n", a, b)

	// Example 2: Convert float64 to int
	var x float64 = 15.5
	var y int = int(x)
	fmt.Printf("x: %f, y: %d\n", x, y)

	// Example 3: Convert int to string
	var num int = 42
	var str string = fmt.Sprintf("%d", num)
	fmt.Printf("num: %d, str: %s\n", num, str)
}
