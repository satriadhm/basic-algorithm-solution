package main

import "fmt"

func main() {
	Binary(34)
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
