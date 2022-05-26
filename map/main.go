package main

import "fmt"

func main() {

	var chiken map[string]int

	chiken = map[string]int{}

	chiken["january"] = 50
	chiken["february"] = 40

	fmt.Println("january", chiken["january"])
	fmt.Println("mei", chiken["mei"])

	// inisiasi nilai map
	var data map[string]int
	data = map[string]int{}
	data["one"] = 1
	fmt.Println(data["one"])

	// map menggunakan for-range
	var chiken1 = map[string]int{
		"januari" : 50,
		"februari" : 40,
		"maret" : 34,
		"april" : 67,
	}

	for key, val := range chiken1 {
		fmt.Println(key, " \t:" , val)
	}

	// menghapus item map
	var chiken2 = map[string]int{"januari" : 50, "february" : 40}

	fmt.Println(len(chiken2))
	fmt.Println(chiken2)

	delete(chiken2, "januari")

	fmt.Println(len(chiken2))
	fmt.Println(chiken2)

	// deteksi item
	var chiken3 = map[string]int{"januari":50, "februari" :40}
	var value, isExist = chiken3["mai"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not found")
	}

	// kombinasi slice dan map
	var chikens4 = []map[string]string{
		map[string]string{"name":"chiken blue", "gender":"male"},
		map[string]string{"name": "chiken red", "gender":"male"},
		map[string]string{"name": "chiken yellow", "gender": "female"},
	}

	for _, chiken5 := range chikens4 {
		fmt.Println(chiken5["gender"], chiken5["name"])
	}

	// deklarasi map terbaru
	var chikens6 = []map[string]string{
		{"name": "chiken blue", "gender": "male"},
		{"name": "chiken red", "gender": "male"},
		{"name": "chiken yelow", "gender": "female"},
	}

	for _, chiken7 := range chikens6 {
		fmt.Println(chiken7["gender"], chiken7["name"])
	}
}