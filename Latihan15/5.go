package main

import "fmt"

func main() {
	var angka, pangkat int
	fmt.Print("Input angka : ")
	fmt.Scan(&angka)

	fmt.Print("Pangkat : ")
	fmt.Scan(&pangkat)

	fmt.Println("Hasilnya adalah", power(angka, pangkat))
}

func power(x, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * power(x, y-1)
	}
}
