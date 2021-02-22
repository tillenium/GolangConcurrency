package main

import (
	"fmt"
	"time"
)
func boring_return(msg string) <- chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {

			//waits for the receiver to be ready to receive the value.
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	return c
}
func main() {

	/**
	A very simple example of using Select. Basically select statement listens to different channels and parallely makes
	the first value available per channel. Incase multiple values are available it pseudo randomly takes one.

	NOTE: The select will execute multiple time in the loop and would proceed with the exit of the program. It doesn't wait!.
	You also cannot depend on the order here.

	*/
	c1 := boring_return("boring!")
	c2 := boring_return("Nah not boring!")
	for i := 0; i < 5; i++ {

		select {
		case v1 := <-c1:
			fmt.Printf("You can %q \n", v1)
		case v2 := <-c2:
			fmt.Printf("You can %q \n", v2)
		default:
			fmt.Println("Not one is ready to communicate")

		}
		//fmt.Printf("You can %q \n", <-c1) //receive just as a value
		//fmt.Printf("You can %q \n", <-c2) //receive just as a value
	}
	fmt.Println("Leaving ")
}
