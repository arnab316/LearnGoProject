package main

import "fmt"

/*
In the Go language, the struct is a collection of data fields with
declared data types. It is similar to a class in other
object-oriented programming languages. Using struct, you can declare and construct your data
types in Golang. Each field within a struct has a name and a type,
and the fields are used to store data.
*/
type Student struct {
	name    string
	roll_no int
	marks   int
}

func main() {
	var stu Student
	// Student deatails
	stu.name = "Arnab Golder"
	stu.roll_no = 5
	stu.marks = 100
	fmt.Println("Student deatails :", stu)
	// Printing student detail using dot (.) operator
	stu.name = "Mark"
	fmt.Println("Name", stu.name)

}
