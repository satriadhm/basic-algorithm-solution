package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	var result float64
	result = math.Sqrt(((a-c)*(a-c) + (b-d)*(b-d)))
	return result
}

func didalam(cx, cy, r, x, y float64) bool {
	var jarakTitik float64
	jarakTitik = jarak(cx, cy, x, y)
	return jarakTitik <= r
}

func main() {
	var cx, cy, r, cx2, cy2, r2, x, y float64
	var kondisi, kondisi2 bool
	fmt.Scanln(&cx, &cy, &r)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)
	kondisi = didalam(cx, cy, r, x, y)
	kondisi2 = didalam(cx2, cy2, r2, x, y)

	if kondisi && kondisi2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if kondisi && !kondisi2 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if !kondisi && kondisi2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
