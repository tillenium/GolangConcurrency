package main

import (
	"fmt"
	"time"
)

func boring_return1(msg string) <- chan string {
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

/**
Redirects the output from multiple channel to one channel.

The information from multiple channels are fetched concurrently ,
so that one channel is not blocking other parallel channel, which was
happening in previous case.
Also known as Generator Channel.
 */
func fanIn(ch1, ch2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-ch1
		}
	}()

	go func() {
		for {
			c <- <-ch2
		}
	}()

	return c
}

func main(){
	//MULTIPLEXING
	c := fanIn(boring_return1("boring!"), boring_return1("Nah not boring!"))
	for i := 0; i < 10; i++ {

		//the main method would wait fo the value to be sent
		fmt.Printf("You can %q \n", <-c) //receive just as a value
	}
	fmt.Println("Leaving ")
}
