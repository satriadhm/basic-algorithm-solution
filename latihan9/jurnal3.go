package main

import "fmt"

func pow(a, b int) int {

	sum := 1
	for i := 1; i <= b; i++ {
		sum *= a
	}
	return sum
}

func konversi(biner int, desimal *int) {
	var lastnum, jumlah, i int
	for biner > 0 {
		lastnum = biner % 10
		jumlah = jumlah + lastnum*pow(2, i)
		biner = biner / 10
		i++
	}
	*desimal = jumlah
}

func main() {
	var bin, des int
	fmt.Scanln(&bin)
	konversi(bin, &des)
	fmt.Println(des)
}
