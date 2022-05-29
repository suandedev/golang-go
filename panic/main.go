package main

import (
	"fmt"
	"strings"
	"errors"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("canot be empty")
	}
	return true, nil
}

func main(){
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		panic(err.Error())
		fmt.Println(err)
	}
}