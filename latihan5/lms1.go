package main

import "fmt"

func main() {
	var x1, y1, z1 float64
	var x2, y2, z2 float64
	var a, b float64
	fmt.Scan(&x1, &y1, &z1)
	fmt.Scan(&x2, &y2, &z2)
	a = (2*(x1))/(x1+y1+z1) + (y1 / 5) + (z1)
	b = (2*(x2))/(x2+y2+z2) + ((y2) / 5) + (z2)
	fmt.Println(a)
	fmt.Println(b)
}
