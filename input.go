package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" //used to convert input to different formats.
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //scanner creation
	fmt.Println("Type Something: ")
	scanner.Scan()                                       //scans the input
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) //assigns input to the text format of what ever was typed in
	fmt.Printf("You will be %d at the end of 2021", 2021-input)
}
