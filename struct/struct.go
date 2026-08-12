package struc

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	/*
		You can also define a struct within a struct. This is called a nested struct.
	*/
	ownerInfo owner
}

type owner struct{
	name string
}

/*
	You can also define a method for a struct. This is called a receiver function.
	This will allow you to access the struct's fields within the function.
*/

func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func willGetThere (e gasEngine, miles uint8) {
	if miles < e.milesLeft() {
		fmt.Println("You will make it!")
	}else {
		fmt.Println("You need to refuel!")
	}
}
func struc() {
	var engine gasEngine
	fmt.Println(engine.mpg, engine.gallons) 
	/*As we haven't defined the values
	The value will be the default value of the datatype*/

	var engine2 gasEngine = gasEngine{mpg: 25, gallons: 10}
	fmt.Println(engine2.mpg, engine2.gallons)

	/*You can also define a struct within a struct. This is called a nested struct.*/
	var engine3 gasEngine = gasEngine {25, 15 , owner{"Erik"}}
	fmt.Println(engine3.mpg, engine3.gallons, engine3.ownerInfo.name)

	/* 
		Define an anonymous struct, which is a struct without a name.
		You can use it if you are defining a struct that you will only use once.
	*/

	var engine4 = struct {
		mpg uint8
		gallons uint8
	}{25,15}
	fmt.Println(engine4.mpg, engine4.gallons)

	/* Calling the method of a struct*/
	fmt.Printf("Total miles left in the tank: %v", engine.milesLeft())
}