package main

/*
Implicit Type Conversion: Golang does
not support implicit type conversion.
You will encounter an error if you attempt
to assign an int variable to a float variable.
*/
import (
	"fmt"
	"strconv"
)

func main() {
	//Declare a variable
	x := 5.6

	// convert x to float64
	y := int(x)

	//printf
	fmt.Println(y)

	var i int = 45
	var s string = strconv.Itoa(i)
	fmt.Println(s)
}
