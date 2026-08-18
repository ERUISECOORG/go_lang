package generics

import "fmt"

func generic(){
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))
}

func sumSlice[T int | float32 | float64] (slice []T) T{
	// sumSlice [T any] (slice []T) bool
	var sum T
	for _, v := range slice{
		sum += v
	}
	return sum
}

//If you are using structs in the parameters, you can't use the generics