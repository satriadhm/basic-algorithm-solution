package main

type v23 struct {
	v2, v3 int
}
type Trec struct {
	v1 int
	vx v23
	v4 int
}

func main()               {}
func banyakNilai(A *Trec) {}

type tabInt [9999]int

func minimum(tab tabInt, n int) int {
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

type arr [9999]int

func total(tab arr, n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum = sum + tab[i]
	}
	return sum
}
func mean(tab arr, n int) int {
	return total(tab, n) / n
}
func maximum(tab tabInt, n int) int {
	var max int
	var k int
	var idx int
	max = tab[0]
	k = 1
	for k < n {
		if max < tab[k] {
			max = tab[k]
			idx = k
		}
		k++
	}
	return idx
}
