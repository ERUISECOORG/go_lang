package functions
import "fmt"
import "errors"

func functions() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 10
	var denominator int = 2
	var quotient, remainder, err = divideInt(numerator, denominator)

	/* 
		If, else if and else statements are used to control the flow of the program.
	*/
	if err != nil {
		fmt.Printf("%s", err.Error())
	}else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", quotient)
	}else{
		fmt.Printf("The result of the integer division is %v with the remainder %v", quotient, remainder)
	}

	/*
		Switch statement can also be used to control the flow of the program.
	*/

	switch {
		case err!=nil:
			fmt.Printf("%s", err.Error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v", quotient)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v", quotient, remainder)
	}

	/*
		You can also use switch statement with a variable.
	*/
	switch remainder {
		case 0:
			fmt.Printf("The result of the integer division is %v", quotient)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v", quotient, remainder)
	}
}

func printMe(printValue string){
	fmt.Println(printValue)
}

/*
	You need to define the datatype of the return value of the function.
	 If you don't, it will return nothing.
	 You can also define multiple return values by separating them with a comma.
*/
func divideInt (numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Denominator cannot be zero")
		return 0, 0, err
	}
	var quotient int = numerator / denominator
	var remainder int = numerator % denominator
	return quotient, remainder, err
}