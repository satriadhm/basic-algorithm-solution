package main

import "fmt"

const mark = '.'
const space = '_'

var lkata int
var ltotal int
var Nkata int

var pita string
var indeks int
var CC byte
var EOP bool

func main() {
	fmt.Scan(&pita)
	ltotal = 0
	Nkata = 0
	startkata()
	for lkata != 0 {
		ltotal = ltotal + lkata
		Nkata++
		advkata()
	}
	if Nkata != 0 {
		fmt.Printf("%.2f", float64(ltotal)/float64(Nkata))
	} else {
		fmt.Println("Pita kosong !")
	}
}

func START() {
	indeks = 0
	CC = pita[indeks]
	EOP = CC == mark
}

func ADV() {
	indeks = indeks + 1
	CC = pita[indeks]
	EOP = CC == mark
}

func ignoreblank() {
	for CC == space {
		ADV()
	}
}

func menghitungpanjang() {
	lkata = 0
	for CC != space && !EOP {
		lkata = lkata + 1
		ADV()
	}
}

func startkata() {
	START()
	ignoreblank()
	menghitungpanjang()
}

func advkata() {
	ignoreblank()
	menghitungpanjang()
}
