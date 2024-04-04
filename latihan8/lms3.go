package main

import "fmt"

func main() {

}

func kabisat(x int) bool {
	var bisat bool
	if x%400 == 0 || (x%4 == 0 && x%100 != 0) {
		bisat = true
	}
	return bisat
}

func mencaryTahun(x int, out *int) {
	var bysat int
	for true {
		fmt.Scan(&bysat)
		if kabisat(bysat) == true {
			*out = bysat
			break
		}

	}
}
