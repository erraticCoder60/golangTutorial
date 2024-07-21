package main

import "fmt"

//global scope
var globalVar string = "This is global scope varibale"

//function scope

func funcScope() {
	funcVar := "This is function scope variable"
	fmt.Println("func scope varibale value is:", funcVar)
	fmt.Println("global scope variable value is:", globalVar)
}

func main() {
	fmt.Println("global scope variable value is:", globalVar)
	fmt.Println("Hello, erratic coder!")
	//Block Scope
	{
		blockVar := 5
		fmt.Println("Block scope variable value is:", blockVar)
		fmt.Println("global scope variable value is:", globalVar)
	}
	funcScope()
}
