package main

import "fmt"

func main() {

	var names [4]string

	names[0] = "made"
	names[1] = "kiran"
	names[2] = "kak dipa"
	names[3] = "ica"

	fmt.Println(names[0], names[1], names[2], names[3])

	// inisiasi nilai awal array

	var fruits = [4]string{"apel", "jeruk", "semangka", "duren"}

	fmt.Println("jumlah element", len(fruits))
	fmt.Println("isi semua element ", fruits) 

	// inisiasi array dengan gaya vertikal
	var fruits2 [4]string

	// horizontal
	// fruits2 = [4]string {"apel", "jeruk", "semangka", "melon"}

	// vertikal
	fruits2 = [4]string {
		"apel",
		"jeruk",
		"mangga",
		"melon",
	}
	fmt.Println("isi semua element ", fruits2) 

	// tanpa jumlah elemen
	var numbers = [...]int{2,3,4,5}

	fmt.Println("datta array \t:", numbers)
	fmt.Println("jumlah array \t:", len(numbers))

	// multidimensi
	var numbers1 = [2][3]int{[3]int{3,2,3}, [3]int{3,4,5}}
	var numbers2 = [2][3]int{{6,7,8}, {9,7,6}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// perulangan array
	var animals = [4]string{"kusing", "sapi", "ayam", "kelinci"}

	for i:= 0; i < len(animals); i++ {
		fmt.Printf("elemen %d : %s\n", i, animals[i])
	}

	// variabel underscore dalam for-range
	var plants = [3]string{"jahe", "kunyit", "lengkuas"}

	for _, plant := range plants {
		fmt.Printf("nama tanaman : %s\n", plant)
	}

	// alokasi elemen array
	var seasonings = make([]string, 2)
	seasonings[0] = "bawang"
	seasonings[1] = "terasi"

	fmt.Println(seasonings)
}