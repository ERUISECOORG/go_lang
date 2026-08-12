package arrays

import (
	"fmt"
	"time"
)

func arrays() {
	var intArr [3]int32 /*This will create an array of three 0s*/
	intArr[0] = 1
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0]) /*This will print the memory location of the element*/

	intArr2 := [...]int32{1,2,3} /*This will initialize the array, inferring the size due to the "[...]"*/
	fmt.Println(intArr2)
}

func slizers() {
	intArr := [...]int32{1,2,3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4,5,6} /*If you don't define the length, you will create a slice*/
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7) /*You can append to a slice*/
	/*
		In the background, it will create a new array with the new size and copy the old values to the new array.
		The new array will be created in a different memory location.
		The capacity of the slice will be doubled when the new array is created.
	*/

	var intSlice2 []int32 = []int32{8,9,10}
	intSlice = append(intSlice, intSlice2...) /*You need to use "..." to append a slice to another slice*/

	var intSlice3 []int32 = make([]int32,3,8) 
	/*You can use "make" to create a slice with a set capacity.
	This will allow it to grow without reallocating memory*/
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8) 
	/* 
	Map is basically a dictionary. It is a collection of key-value pairs.
	"[string]" defines the key values type
	"uint8" defines the value type
	*/
	fmt.Println(myMap)

	var MyMap2 = map[string]uint8{"value1": 1, "value2": 2} /*You can also initialize a map with values*/
	fmt.Println(MyMap2)

	fmt.Println(MyMap2["value1"]) /*You can access the value of a key by using the key*/
	fmt.Println(MyMap2["value3"]) /*If the key doesn't exist, it will return the default of the value "Type"*/

	var age, ok = MyMap2["Jason"] /*You can also check if the key exists by using the second return value*/
	if ok != true {
		fmt.Println("Key doesn't exist")
	} else
	{
		println("Key exists with value: ", age)
	}

	/*
		You can use the "range" keyword to iterate over a map.
		It will return the key and value of each element in the map.
	*/
	for name := range MyMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	for name, age := range MyMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	/* 
		You can also use it to iterate on arrays
		This will return the index and value of each element in the array.
	*/
	for index, value := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}

	/*
		While loop example
	*/
	var i int = 0
	for i < 10{
		i++
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/*
		Quick time function usage example
	*/

	var startTime = time.Now()
	for i := 0; i < 1000000; i++ {
	}
	var delay = time.Since(startTime)

	fmt.Println(delay)
}