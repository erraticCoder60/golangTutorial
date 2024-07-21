package main

import "fmt"

func pointerFunc(pointer_var *int) {
	*pointer_var = 20
}

func main() {
	fmt.Println("Hello Erratic Coder")

	//Declaration
	var ptr *int

	//Initiatilization
	var num int = 10
	ptr = &num

	//Dereferencing
	fmt.Println(*ptr)

	pointerFunc(ptr)

	fmt.Println(*ptr)

}
