package main

import "fmt"
import "strconv"

func main() {
	var str = "123"
	var num, err = strconv.Atoi(str)

	var num2, err2 = strconv.ParseInt(str, 10, 64)

	if err == nil {
		fmt.Println(num)
	}

	if err2 == nil {
		fmt.Println(num2)
	}
}