package main

type tabInt [9999]int

func NilaiMinimum(tab tabInt, n int) int {
	var min int
	var k int
	var idx int
	min = tab[0]
	k = 1
	for k < n {
		if min > tab[k] {
			min = tab[k]
			idx = k
		}
		k++
	}
	return idx

}
