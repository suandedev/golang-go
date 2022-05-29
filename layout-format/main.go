package main

import "fmt"

type student struct {
	name string
	height  float64
	age int32
	isGraduated bool
	hobbies []string
}

func main() {
	var data = student{
		name: "made",
		height: 1.75,
		age: 20,
		isGraduated: true,
		hobbies: []string{"coding", "reading"},
	}

	fmt.Printf("%b\n", data.age) // biner 1010
	fmt.Printf("%d\n", data.age) //string
	fmt.Printf("%e\n", data.height) //sientific
	fmt.Printf("%s\n", data.name) //string
	fmt.Printf("%t\n", data.isGraduated) //boolean
	fmt.Printf("%T\n", data.name)
	// string

	fmt.Printf("%T\n", data.height)
	// float64

	fmt.Printf("%T\n", data.age)
	// int32

	fmt.Printf("%T\n", data.isGraduated)
	// bool

	fmt.Printf("%T\n", data.hobbies)
	// []string

	fmt.Printf("%v\n", data)
// {wick 182.5 26 false [eating sleeping]}
fmt.Printf("%+v\n", data)
// {name:wick height:182.5 age:26 isGraduated:false hobbies:[eating sleeping]}
fmt.Printf("%#v\n", data)
// main.student{name:"wick", height:182.5, age:26, isGraduated:false, hobbies:[]string{"eating", "sleeping"}}
}