package main

import "fmt"

func main4() {
	var a, b float64
	var c bool
	for selesai := false; !selesai; {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&a, &b)
		c = a-b >= 9 || b-a >= 9
		fmt.Println("Sepeda motor pak Andi akan oleng: ", c)
		selesai = a+b >= 150 || a <= -1 || b <= -1
	}
}
