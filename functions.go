package main

import "fmt"

func multipleNumbers(num1, num2 int) int {
	mul := num1 * num2
	return mul
}

func main() {
	//function call
	res := multipleNumbers(5, 4)
	fmt.Print(res)
}
