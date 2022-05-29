package main

import (
	"fmt"
	"strings"
	"errors"
)

func catch() {
	if r := recover(); r!= nil {
		fmt.Println("error closure", r)
	} else {
		fmt.Println("running perfect")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("canot be empty")
	}
	return true, nil
}

func main() {
	defer catch()

	var name string
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("error")
	}
}