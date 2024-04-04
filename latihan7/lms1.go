package main

import "fmt"

func main() {
	var n int
	var berat, tinggi, imt float64

	fmt.Scan(&n)
	i := 0
	for i < n {
		fmt.Scan(&berat, &tinggi)
		imt = berat / (tinggi * tinggi)
		if imt < 18.5 {
			fmt.Print("Kurus")
		} else if imt >= 18.5 && imt < 25 {
			fmt.Print("Normal")
		} else if imt >= 25 && imt <= 30 {
			fmt.Print("Gemuk")
		} else if imt > 30 {
			fmt.Print("Obesitas")
		}
		i++

	}
}
