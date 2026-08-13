package main

import "fmt"

func main(){
	/*
		Pointers are variables that store memory locations
		default value for pointers is nil
	*/
	/* 
		var p *int32 intializes the pointer
		new(int32) gives the pointer a memory location to point into
	*/
	var p *int32 = new(int32) 
	var i int32
	*p = 10
	fmt.Println(*p)
	fmt.Println(i)

	/* Pointers can also reference to the memory space of a variable*/
	p = &i
	/* 
		If you change the pointer value, you change the value stored in that memory space
		As a variable points to that memory space, the variable value will also change
	*/
	*p = 1
	fmt.Println(*p)
	fmt.Println(i)

	/* This is different from when you assign the value of a variable into another */
	/* That copies the value from one memory location into a new memory location */
	var k int32 = 2
	i = k
}