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

//Quits after 10 seconds here
func quit_channel() <- chan bool {
	quit := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		quit <- true
	}()
	return quit
}

func main() {

	/**
	A timeout strategy can be used inside select

	*/
	c1 := boring_return("boring!")
	c2 := boring_return("Nah not boring!")
	quit := quit_channel()
	//this is timeout whole service.
	overall_timeout := time.After(5 * time.Second)
	for {

		select {
		case v1 := <-c1:
			fmt.Printf("You can %q \n", v1)
		case v2 := <-c2:
			fmt.Printf("You can %q \n", v2)

			//THis is going to wait for all the channel to supply the value with in this time.
		case <- time.After(100 * time.Millisecond):
			fmt.Println("Timeout of this call.")
			//return
		case <- overall_timeout:
			fmt.Println("The service got timeout. Closing")
			//return
		case <- quit:
			fmt.Println("Using quit channel to quit")
			return

		}
		//fmt.Printf("You can %q \n", <-c1) //receive just as a value
		//fmt.Printf("You can %q \n", <-c2) //receive just as a value
	}
	fmt.Println("Leaving ")
}
