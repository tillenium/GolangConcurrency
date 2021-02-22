package main

import (
	"fmt"
	"time"
)

/**
Channel is basically a connection between two go routines.
*/

func boring_channel(msg string, c chan string) {
	for i := 0; ; i++ {

		//waits for the receiver to be ready to receive the value.
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(1000 * time.Millisecond)
	}
}
func main() {
	c := make(chan string)
	go boring_channel("boring!", c)
	for i := 0; i < 5; i++ {

		//the main method would wait fo the value to be sent
		fmt.Printf("You can %q \n", <-c) //receive just as a value
	}
	fmt.Println("Leaving ")
	/**
	What's happening with channel that the main method is the controller here.
	It waits on the channel for 5 values and then just exits.
	*/
}
