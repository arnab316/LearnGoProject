package main

import "fmt"

func main() {

	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	//for shorter syntax
	arrayName := [3]string{"Apple", "Kiwi", "Orange"}
	fmt.Println(arrayName)
	//fmt.Println(arrayName[0])
	fmt.Println(len(arrayName))
	fmt.Print(cap(arrayName))
}
