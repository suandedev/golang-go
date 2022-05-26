package main

import "fmt"

type siswa struct {
	name string
	grade int
}

func (s siswa) changeName1(name string) {
	fmt.Println("on change name1, name change to", name )
	s.name = name
}

func (s *siswa) changeName2(name string) {
	fmt.Println("on chage name2, name chaged on", name)
	s.name = name
}

func main() {
	var s1 = siswa{"made suande", 21}
	fmt.Println("s1 before ", s1.name)

	s1.changeName1("komang")
	fmt.Println("s1 after chage name 1", s1.name)

	s1.changeName2("wayan")
	fmt.Println("s1 affter change name 2", s1.name)
}