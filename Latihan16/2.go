package main

import "fmt"

var CC byte
var pita string
var EOP bool
var indeks int

func main() {
	fmt.Scan(&pita)
	START()

	var sum int
	var stat int
	sum = 0
	stat = 1
	for !EOP {
		if stat == 1 {
			if CC == 'F' {
				stat = 2
			}
		} else if stat == 2 {
			if CC == 'I' {
				stat = 3
			} else {
				stat = 1
			}
		} else if stat == 3 {
			if CC == 'F' {
				sum++
			}
			stat = 1
		}
		ADV()
	}
	fmt.Println(sum)
}

func ADV() {
	indeks = indeks + 1
	CC = pita[indeks]
	EOP = CC == '.'
}
func START() {
	indeks = 0
	CC = pita[indeks]
	EOP = CC == '.'
}
