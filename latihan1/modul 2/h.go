package main

import "fmt"

func hitungVolume(r int, t int) float64 {
	var pi float64 = 3.14
	return float64(r) * float64(r) * pi * float64(t)
}

func main() {
	var r, t int
	fmt.Scan(&r)
	fmt.Scan(&t)
	fmt.Println(hitungVolume(r, t))

}
