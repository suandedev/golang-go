package main

import "fmt"

func main() {
	var student struct {
		person
		grade int
	}

	student.person = person{"made", 20}
	student.grade = 2
}