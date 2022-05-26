package library

import f "fmt"

var Student = struct {
	Name string
	Grade int
}{}

func init() {
	Student.Name = "made suande"
	Student.Grade = 21
	 f.Println("--> library/library.go imported")
}

