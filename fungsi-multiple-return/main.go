package main

import "fmt"
import "math"

func calculated(d float64) (area float64, circumference float64) {
	// hitung luas
	area = math.Pi * math.Pow(d/ 2, 2)

	// hitung keliling
	circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

func main() {

	var diameter float64 = 15

	var area, circumference = calculated(diameter)

	fmt.Println("luas keliling", area)
	fmt.Println("keliling lingkaran", circumference)
}