package main

import f "fmt"
import "math"

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

type persegi struct {
	sisi float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	var bagunDatar hitung

	bagunDatar = persegi{10.0}
	f.Println("==== persegi")
	f.Println("luas :", bagunDatar.luas())
	f.Println("keliling :", bagunDatar.keliling())

	bagunDatar = lingkaran{14.0}
	f.Println("=== lingkaran")
	f.Println("luas", bagunDatar.luas())
	f.Println("keliling", bagunDatar.keliling())
	f.Println("jari jari", bagunDatar.(lingkaran).jariJari())

}



