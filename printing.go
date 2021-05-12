package main

import "fmt"

func main() {
	fmt.Printf("Hello %T %v\n", 10, 10)                               //T is used to print the data type, v is used to print value
	var x string = fmt.Sprintf("Hello from the Sprint %T %v", 10, 10) // Sprintf allows us to create a format output and store it in a variable
	fmt.Println(x)

	//base formatting

	fmt.Printf("Base 2 of 12340: %b\n", 12340) //Prints out the base 2 representation of the number - binary
	fmt.Printf("Base 8 of 3435: %o\n", 3435)   // Prints out the base 8 representation of the number - octal
	fmt.Printf("Base 10 of 8371: %d\n", 8371)  // Prints out the base 10 representation of the number - decimal
	fmt.Printf("Base 16 of 1535: %x\n", 1535)  // Prints out the base 16 representation of the number - hex
	fmt.Println("--------------------------------------------------------------------")

	//floating points
	fmt.Printf("Scientific notation of %s: %e\n", "312.4321111111111", 312.4321111111111)
	fmt.Printf("Floating point w/o notation of %s: %f\n", "312.4321111111111", 312.4321111111111)
	fmt.Printf("Large exponent of %s: %g\n", "312.4321111111111", 312.4321111111111)
	fmt.Println("--------------------------------------------------------------------")
	//strings
	fmt.Printf("Strings use the %%s format: %s\n", "Alex")                //Prints out regular string
	fmt.Printf("String with quotations is achieved by %%q: %q\n", "Alex") // Surrounds string with quotes
	fmt.Println("--------------------------------------------------------------------")
	//Width and Precision
	fmt.Println("Widths include the # of indexs from input. EX: 9 width of 10 is --------10 (8 spaces with 2 int symbols)")
	fmt.Printf("Default width, default precision of %s: %f\n", "3.432", 3.432)
	fmt.Printf("Width of 9, default precision of %s: %9f\n", "3.43254", 3.43254)
	fmt.Printf("Default width, 2 precision of %s: %.2f\n", "3.43254", 3.43254)
	fmt.Printf("Width of 9, 2 precision of %s: %9.2f\n", "3.43254", 3.43254)
	fmt.Printf("Width of 9, 0 precision of %s: %9.f\n", "3.43254", 3.43254)
	fmt.Printf("Width of -9, default precision of %s: %-9s is cool\n", "Alex", "Alex")
	//Padding
	fmt.Println("--------------------------------------------------------------------")
	fmt.Printf("PADDING Width of 9, default precision of %s: %09d\n", "3", 3)

}
