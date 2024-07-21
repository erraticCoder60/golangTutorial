package main

import "fmt"

func main() {
	fmt.Println("Welcome on erratic coder!")

	//Array

	// var arr [3]int
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3

	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("array is:", arr)

	//Slices

	slice := arr[2:4]

	fmt.Println("slice is:", slice)

	fmt.Println("length of slice is:", len(slice))

	fmt.Println("Capacity of slice is:", cap(slice))

	slice = append(slice, 10)

	fmt.Println("appended slice is:", slice)

	//make function

	slice1 := make([]int, 3, 6)
	slice1[0] = 100
	slice1[1] = 200
	slice1[2] = 300

	slice1 = append(slice1, 400)

	fmt.Println("Slice with make function is:", slice1)

	fmt.Println("length of Slice with make function is:", len(slice1))

	fmt.Println("Capacity of Slice with make function is:", cap(slice1))

}
