package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	var intVar int
	fmt.Println(intVar)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello \n World"
	fmt.Println(myString)
	var concatString string = "Hello" + " " + "World"
	fmt.Println(concatString)

	fmt.Println(len("a"))
	print(utf8.RuneCountInString("abcde"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	/* Default Values
	   For int, uint, float and rune is 0
	   For string is "" emtpy string
	   And for bool is 0
	*/

	undeclared_var := "text"
	/* 
		undeclared_var := foo().
		It is useful when you are not certain on what the output of a function will be.
		Otherwise you should be good enforcing the type.
	*/
	fmt.Println(undeclared_var)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Println(myConst)

	const pi float32 = 3.1415
}