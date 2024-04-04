package main

import "fmt"

func main() {
	var r int
	var luas float64
	fmt.Scan(&r)
	luas = 22 / 7 * float64(r)
	fmt.Println("Luas lingkaran dengan jari-jari =", r, "adalah", luas)

}
