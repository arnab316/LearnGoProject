package main

import "fmt"

func main() {
	primes := []int{1, 2, 3, 4, 5}
	var slice []int = primes[1:3]
	fmt.Println(slice)
}
