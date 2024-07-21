package main

import (
	"fmt"
	"unicode"
)

func main() {

	str := "Hello123"

	for _, value := range str {
		if unicode.IsDigit(value) {
			fmt.Println("Character is digit")
		} else {
			fmt.Println("Character is letter")
		}
	}
}
