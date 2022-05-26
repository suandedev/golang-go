package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	var s1 = struct {
		person
		grade int
	}{}

	s1.person = person{"made", 21}
	s1.grade = 2

	fmt.Println("name :", s1.name)
	fmt.Println("age :", s1.age)
	fmt.Println("grade :", s1.grade)
}