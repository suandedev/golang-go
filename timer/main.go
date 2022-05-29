package main

import "fmt"
import "time"

func main() {
	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("affter 4 second")
}