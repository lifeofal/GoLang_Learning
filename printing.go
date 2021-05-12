package main

import "fmt"

func main() {
	fmt.Printf("Hello %T %v\n", 10, 10)                               //T is used to print the data type, v is used to print value
	var x string = fmt.Sprintf("Hello from the Sprint %T %v", 10, 10) // Sprintf allows us to create a format output and store it in a variable
	fmt.Println(x)
}
