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
}