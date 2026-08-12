package strings

import "fmt"

func strings() {
	var myString = "résumé"
	var indexed = myString[0]
	/*
		Every element in a string is an array of bytes.
		So when you index a string, you will get the byte value of the character at that index.
	*/
	fmt.Println(indexed)
	fmt.Printf("%v, %T", indexed, indexed)

	var myRune = 'a'
	fmt.Println(myRune) /* A rune is an alias for int32, and it represents a Unicode code point. */
}