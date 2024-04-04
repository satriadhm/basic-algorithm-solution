package main

import "fmt"

func main5() {
	var k float64
	var fk, fk0 float64
	fmt.Print("Nilai K =")
	fmt.Scan(&k)
	fk = ((4*float64(k) + 2) * (4*float64(k) + 2))
	fk0 = ((4*float64(k) + 1) * (4*float64(k) + 3))
	fmt.Println("Nilai f(k) :", fk/fk0)
}
