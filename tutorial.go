package main

import "fmt"

func main() {
	var x int = 5
	z := 10 //implicit

	//z := 10 is equivalent to var z int = 10. GoLang can define the data type automatically.
	//vars cant be a keyword and cannot start with a number or contain a special char.

	number := 500 //implicitly making 'number' an int variable
	number = 2    //assigning 'number' to a new value *MUST BE THE SAME DATA TYPE AS VARIABLE CREATION*

	fmt.Println("Hello world", z, x, number)

}
