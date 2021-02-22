package main

import (
	"fmt"
	"time"
)


/**
Redirects the output from multiple channel to one channel.

Implementing it using swtich and switch takes care of multiplexing.
*/
func fanInUsingSwtch(ch1, ch2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
				case s := <-ch1:
					c <- s
				case s := <-ch2:
					c <- s
			}
		}
	}()

	return c
}

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
	//MULTIPLEXING
	c := fanInUsingSwtch(boring_return("boring!"), boring_return("Nah not boring!"))
	for i := 0; i < 10; i++ {

		//the main method would wait fo the value to be sent
		fmt.Printf("You can %q \n", <-c) //receive just as a value
	}
	fmt.Println("Leaving ")
}
