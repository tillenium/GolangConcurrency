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
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	return c
}

func main(){

	/**
	Can be used to create multiple services which run in parallel and you can
	wait for the values to recieved in sync.

	NOTE: Though the second will be waiting for the first value as it will be blocked for the wait.
	 */
	c1 := boring_return("boring!")
	c2 := boring_return("Nah not boring!")
	for i := 0; i < 5; i++ {

		//the main method would wait fo the value to be sent
		fmt.Printf("You can %q \n", <-c1) //receive just as a value
		fmt.Printf("You can %q \n", <-c2) //receive just as a value
	}
	fmt.Println("Leaving ")
}
