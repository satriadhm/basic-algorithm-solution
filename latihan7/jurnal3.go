package main

import "fmt"

func main() {
	var n, sks, jumlah, totalSks int
	var nilai string
	var ipSemester float64
	fmt.Scan(&n)
	i := 0
	for i < n {
		fmt.Scan(&nilai, &sks)
		for (nilai == "A" && nilai == "B" && nilai == "C" && nilai == "D" && nilai == "E") && sks > 0 {
			fmt.Scan(&nilai, &sks)
		}
		totalSks += sks
		if nilai == "A" {
			jumlah += 4 * sks
		} else if nilai == "B" {
			jumlah += 3 * sks
		} else if nilai == "C" {
			jumlah += 2 * sks
		} else if nilai == "D" {
			jumlah += 1 * sks
		}
		i++
	}
	ipSemester = float64(jumlah) / float64(totalSks)

	fmt.Print(ipSemester)

}
