package main

import "fmt"
import "reflect"

func main() {
	var number = 23
	var reflectvalue = reflect.ValueOf(number)

	fmt.Println("tipe data ", reflectvalue.Type())

	if reflectvalue.Kind() == reflect.Int {
		fmt.Println("nilai var :", reflectvalue.Int())
	}

	// interface

	var number2 = 23
	var reflectvalue2 = reflect.ValueOf(number2)

	fmt.Println("tipe data ", reflectvalue2.Type())
	fmt.Println("nulai variabel", reflectvalue2.Interface())
}