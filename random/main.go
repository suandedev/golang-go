package main

import "fmt"
import "math/rand"
import "time"



func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("random 1:", rand.Int())
	fmt.Println("random 2:", rand.Int())
	fmt.Println("random 3:", rand.Int())

}