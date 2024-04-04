package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n && i < NMAX; i++ {
		fmt.Scan(&data[i])
	}
}
func posisi(n, k int) int {
	for i := 0; i < n && i < NMAX; i++ {
		if data[i] == k {
			return i
		}
	}
	return -1
}
func main() {
	var n, k, hasil int
	fmt.Scan(&n, &k)
	isiArray(n)
	hasil = posisi(n, k)
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
