package main

import "fmt"

var pita string
var CC byte
var EOP bool
var indeks int

func main() {
	dash := 0
	fmt.Scan(&pita)
	START()
	kelas := pita[:2] == "SE" || pita[:2] == "IF" || pita[:2] == "IT"
	indeks = 2
	var nidigit bool = true
	for !EOP {
		ADV()
		if pita[indeks-1] == '-' {
			dash++
			nidigit = nidigit && DIGIT()
		}
	}
	fmt.Print(kelas && nidigit && dash == 2)
}

func START() {
	indeks = 0
	CC = pita[indeks]
	EOP = CC == '.'
}

func ADV() {
	indeks++
	CC = pita[indeks]
	EOP = CC == '.'
}

func DIGIT() bool {
	p := 0
	for CC != '-' && CC != '.' {
		if CC > 47 && CC < 58 {
			p++
		}
		ADV()
	}
	return p > 0
}
