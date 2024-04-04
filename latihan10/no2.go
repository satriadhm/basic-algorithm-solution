package main

import "fmt"

func main() {
	var n int
	var a1, a2, a3, b1, b2, b3 int
	var sumA, sumB int
	fmt.Scan(&n)
	sumA = 0
	sumB = 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&a1, &a2, &a3, &b1, &b2, &b3)
		sumA += a1 + a2 + a3
		sumB += b1 + b2 + b3
	}
	fmt.Println("Petinju A: ", sumA, "dan petinju B: ", sumB)
	if sumA > sumB {
		fmt.Println("Pemenang adalah petinju A")
	} else if sumB > sumA {
		fmt.Println("Pemenang adalah petinju B")
	}
}
