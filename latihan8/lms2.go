package main

import "fmt"

func main() {
	var x int = 2
	var y int = 3
	var z int = 5
	fmt.Printf("%.2f", mergeFG(x, y, z))
}
func findF(x, y, z int, hasil *float64) {
	*hasil = float64(2*x) / float64(x+y+z)
}

func searchG(x, y int, hasil *float64) {
	*hasil = float64(x+5*y) / 5
}

func mergeFG(x, y, z int) float64 {
	var hasil1, hasil2 float64
	findF(x, y, z, &hasil1)
	searchG(y, z, &hasil2)
	sum1 := hasil1 + hasil2
	return sum1
}
