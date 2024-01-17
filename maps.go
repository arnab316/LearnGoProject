package main

import "fmt"

/*
In the Go language, the map data structure is a
collection of key-value pairs,
where each key is unique.
*/
func main() {
	prog := map[string]int{"java": 30, "javascript": 90, "Golang": 100}
	//fmt.Print(prog)
	fmt.Println(len(prog))
}
