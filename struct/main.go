package main

import "fmt"

// deklarasi
type student struct {
	name string
	grade int
}

func main() {
// penerapan struct
	var s1 student
	s1.name = "made suande"
	s1.grade = 2

	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)

	// inisiasi object struct
	var s2 = student{"komang", 2}
	var s3 = student{name : "wayan"}
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)

	// variabel object pointer
	var s4 = student{name: "komang wi", grade: 4}

	var s5 *student = &s4
	fmt.Println("student 4, name :", s4.name)
	fmt.Println("student 5, name :", s5.name)

	s5.name = "wayan wi"

	fmt.Println("student 4, name :", s4.name)
	fmt.Println("student 5, name :", s5.name)
}