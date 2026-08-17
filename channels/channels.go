package channels

import "fmt"
import "time"


func channels() {
	/* This will cause a dead lock error, because once you add a value to a channel it waits for something to pull that value
	   The reason for this is that channels are meant to be used along with routines

	var c = make(chan int) make(chan <type of value that channel will have>)*
	c <- 1                 To add a value to a channel, we use "<-"*
	var i = <-c            This line will pop the value out of the channel and add it as value to the variable*
	fmt.Println(i)
|	*/

	var c = make(chan int, 5) 
	/*
		the "5" indicates the room that the channel will have for values
		this allows the routine to store 5 values in the channel
		without having to wait for the main function to clear the channel before adding a new value.
	*/
	go process(c)
	for i:= range c{
		/*fmt.Println(<-c) -- This line will wait for a value to be set in the channel on the function*/
		fmt.Println(i) 
		time.Sleep(time.Second*1)
	}
}

func process(c chan int){
	 /*
	 	This line will make the channel close once the function is completed, preventing deadlock
	*/
		defer close(c)
	/*
		fmt.Println(<-c) is waiting for a value in the channel or for it to be closed
		This doesn't apply for the "for i:=range c"
	*/
	for i:=range 5{
		c <- i
	}
	fmt.Println("Exiting Process")
}