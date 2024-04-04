package main

import "fmt"

func main() {
	var dec int
	fmt.Scan(&dec)
	x := binary(dec)
	for i := len(x) - 1; i > 0; i-- {
		fmt.Print(x[i])
	}
	Binary(5)
}

func binary(angk int) []int {
	var slc []int
	for angk > 0 {
		mod := angk % 2
		slc = append(slc, mod)
		angk /= 2
	}
	return slc
}

func Binary(angk int) {
	var slc []int
	for angk > 0 {
		mod := angk % 2
		slc = append(slc, mod)
		angk /= 2
		for i := len(slc) - 1; i > 0; i-- {
			fmt.Print(slc[i])
		}
	}

}
