package main

import "fmt"
import "strings"

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each:= range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string {"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)

	fmt.Println("filter ada huruf \"5\"\t:", dataLength5)
	// filter julah huruf "5"
}