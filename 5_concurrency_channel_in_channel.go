package main

import (
	"fmt"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func boring_return2(msg string, t time.Duration) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {

			//waits for the receiver to be ready to receive the value.
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(t * time.Millisecond)
			<-waitForIt
		}
	}()

	return c
}

func main() {

	/**
	Here we would be using WaitforIt channel inside channel to create the order of all by the channel.
	This is similar to the example 3 just using channel inside channel.

	Here the channels have different wait time: 0.5sec and 1 sec
	*/
	c1 := boring_return2("boring!", 500)
	c2 := boring_return2("Nah not boring!", 1000)
	for i := 0; i < 5; i++ {

		//the main method would wait fo the value to be sent
		msg1 := <-c1
		fmt.Printf("You can %q \n", msg1.str) //receive just as a value

		msg2 := <-c2
		fmt.Printf("You can %q \n", msg2.str) //receive just as a value

		/**
		The main program is giving the head go to the two go routines to be able to continue execution.

		Incase we remove any statement there would be deadlock.
		This is another example of channel in the channel.
		*/
		msg1.wait <- true
		msg2.wait <- true

	}
	fmt.Println("Leaving ")
}
