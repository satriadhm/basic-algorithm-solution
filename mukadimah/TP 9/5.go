package main

type arr [9999]int

func NilaiRerata(tab arr, n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum = sum + tab[i]
	}
	return sum / n
}
