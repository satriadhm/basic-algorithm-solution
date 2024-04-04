package main

type arr [10000]int

func CariCepat(tab arr, n, x int) bool {
	var left, right, mid int
	left = 1
	right = n
	mid = (left + right) / 2
	for left <= right && tab[mid] != x {
		if x < tab[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return mid > 0 && tab[mid] == x
}
