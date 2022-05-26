package main

import "fmt"

func main()  {
	//variabel beserta tipe data
	// var firstNAme string = "made"
	// var lastName string = "suande"
	// fmt.Printf("Hello %s %s\n", firstNAme, lastName)

	//deklarasi tanpa tipe data
	var firstName string = "made"
	lastName := "suande"
	fmt.Printf("Hello %s %s\n", firstName, lastName)

	var hewan2 = "kucing"
	var hewan = "sapi"
	fmt.Println(hewan2, hewan)

	//multi variabel
	var kucing, ayam, aning string
	kucing, ayam, aning = "bela", "boiler", "kiki"
	fmt.Println(kucing, ayam, aning)
}