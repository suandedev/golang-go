package main

import "fmt"

func main() {

	type person struct {
		name string
		age int
	}

	var allStudents = []person{
		{name: "made", age: 23},
		{name: "made", age: 22},
		{name: "made", age: 21},
		{name: "made", age: 20},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

}