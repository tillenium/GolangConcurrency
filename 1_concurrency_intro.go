package main

import (
	"fmt"
	"time"
)

func boring() {
	for i := 0; ; i++ {
		fmt.Println("boring", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	/**
	1. Will print the number continuously
	*/
	//boring()

	//------------------

	/**
	2. If we run the function in go routine, it won't run.
	By default when the main function ends so does the go routine.
	*/
	go boring()

	/**
	3. Needs this in order to wait for the go routine to run and main
	would wait.

	NOTE: The main function would wait only 1 sec for the boring to run.
	The boring is an infinite loop, though when main method exists so does,
	the go routine.

	So overall there won't be any dangling go routing running in the
	background.
	*/
	time.Sleep(1000 * time.Millisecond)

	/**
	4. Something about GO:
	- Go has it's own stack so we can just keep on spinning off GO routines.
	There are very cheap. They are like extremely cheap threads.

	At runtime they are multiplexed on the threads.
	 */
}
