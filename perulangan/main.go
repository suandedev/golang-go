package main

import "fmt"

func main() {

	// perulangan hanya kondisi
	var i = 0

	for i < 5 {
		fmt.Println("angka", i)
		i++
	}

	//perualngan tanpa argument
	var i1 = 0

	for {
		fmt.Println("angka", i1)

		i1++
		if i1 == 5 {
			break
		}
	}

	// perulangan break continune
	for i2 := 1; i2 <= 10; i2++ {
		if i2 % 2 == 1 {
			continue
		}

		if i2 > 8 {
			break
		}

		fmt.Println("angka", i2)
	}

	// perulanan bersarang
	for i3 := 0; i3 < 5; i3++ {
		for j := i3; j < 5; j++ {
			fmt.Println(j, " ")
		}
		fmt.Println()
	}

	// label dalam perulangan
	outerLoop:
	for i4 :=0; i4 <5; i4++ {
		for j := 0; j < 5; j++ {
			if i4 == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i4, "][", j, "][", "\n")
		}
	}

}