package main

import "fmt"

func main() {
	day := "Thursday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday!")
		fallthrough
	case "Tuesday":
		fmt.Println("It's Tuesday!")
		fallthrough
	case "Wednesday":
		fmt.Println("It's Wednesday!")
		fallthrough
	case "Thursday":
		fmt.Println("It's Thursday!")
		fallthrough
	case "Friday":
		fmt.Println("It's Friday!")
		fallthrough
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
		fallthrough
	default:
		fmt.Println("Invalid day!")
	}
}
