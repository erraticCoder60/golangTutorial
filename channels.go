package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- string) {
	fmt.Println("This is sender goroutine")
	ch <- "Hello, erratic coder!" // Send string message into the channel

	//Send various values
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Sending", i, "to the channel")
	// 	ch <- i // Send the value to the channel
	// 	fmt.Println("Sent", i, "to the channel")
	// }
}

func receiver(ch <-chan string) {
	fmt.Println("This is receiver goroutine")
	for value := range ch {
		fmt.Println("Received", value, "from the channel")
		time.Sleep(time.Second) // Simulate some work
	}
}

func main() {
	// // Create a un-bufferred channels for communication
	// messageChannel := make(chan string)

	// Create a bufferred channels for communication
	messageChannel := make(chan string, 3)

	// Start a goroutine to send a message
	go sender(messageChannel)

	// Start another goroutine to receive the message
	go receiver(messageChannel)

	// Wait for a keystroke before exiting
	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}
