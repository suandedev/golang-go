package main

import "fmt"
import "time"

func main() {
	var time1 = time.Now()
	fmt.Printf("time %v\n", time1)

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time 2 %v\n", time2)
}