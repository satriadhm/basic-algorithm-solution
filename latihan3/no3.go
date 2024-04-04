package main

import "fmt"

func main3() {
	var a, b float64
	for selesai := false; !selesai; {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&a, &b)
		selesai = a >= 9 || b >= 9
	}
}
