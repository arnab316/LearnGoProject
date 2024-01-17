package main

import (
	"fmt"
)

/*
Pointers in the Go programming
language or Golang is a variable that holds the
memory address of another variable.
*/
/*
The & operator is used to get a variable's
memory address, while the *
operator is used to dereference a pointer and access
the value stored at the memory address
*/

func main() {
	var val int = 80 // Variable declaration
	var ptr *int     // Pointer variable declaration
	ptr = &val
	fmt.Print(ptr)
}
