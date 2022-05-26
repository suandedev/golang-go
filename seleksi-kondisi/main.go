package main

import "fmt"

func main() {
	var point = 8

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Println("Tidak lulus")
	}

	//temporary if
	var point1 = 8840.0

	if percent := point1 / 100; percent > 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s nod bad\n", percent, "%")
	}

	//menggunakan switch
	var point2 = 6
	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// kurung kurawal
	var point3 = 6

	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be batter")
		}
	}

	// switch gaya if
	var point4 = 6

	switch {
	case point4 == 8:
		fmt.Println("perfect")
	case (point4 < 8) && (point4 > 3):
		fmt.Println("awesome4")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be batter")
		}
	}

	// fallthrough
	var point5 = 6

	switch{
	case point5 ==8:
		fmt.Println("perfect")
	case (point5 < 8) && (point5 > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be batter")
		}
	}

	// kodisi bersarang
	var point6 = 10

	if point6 > 7 {
		switch point6 {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if point6 == 5 {
			fmt.Println("nor bad")
		} else if point6 == 3 {
			fmt.Println("kepp trying")
		} else {
			fmt.Println("you can do it")
			if point6 == 0 {
				fmt.Println("try harder")
			}
		}
	}
}