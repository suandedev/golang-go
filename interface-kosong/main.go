package main

import "fmt"

func main() {
	var person = []map[string]interface{}{
		{"name" : "wick", "age" : 21},
		{"name" : "wick", "age" : 21},
		{"name" : "wick", "age" : 21},
	}

	for _, each:= range person {
		fmt.Println(each["name"], "age is", each["age"])
	}
}