package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("Number (address) :", &numberA)

	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("Number (address) :", numberB)


	// parameter pointer
	var number = 4
	fmt.Println("before :", number)

	change(&number, 10)
	fmt.Println("affter :", number)

	
}

func change(original *int, value int) {
	*original = value
}