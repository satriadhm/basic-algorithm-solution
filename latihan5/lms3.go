package main

import (
	"fmt"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println("Bola 1 terberat ? ", a > b && a > c && a > d)
	fmt.Println("Bola 2 terberat ? ", b > c && b > a && b > d)
	fmt.Println("Bola 3 terberat ? ", c > a && c > b && c > d)
	fmt.Println("Bola 4 terberat ? ", d > c && d > a && d > b)
}
