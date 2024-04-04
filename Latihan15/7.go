package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func main() {
	var A arrInt
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	selectionSort(&A, N)
	for i := 0; i < N; i++ {
		fmt.Printf("%v ", A[i])
	}
}

func selectionSort(T *arrInt, n int) {
	//fungsi master dari selection sort
	sort(T, n, 1)
}

func sort(T *arrInt, n, pass int) {
	// fungsi rekursif untuk selection sort
	var imin, temp int
	if pass <= n-1 {
		imin = min(*T, pass, pass-1, n)
		temp = T[pass-1]
		T[pass-1] = T[imin]
		T[imin] = temp
		sort(T, n, pass+1)
	}
}

func min(T arrInt, i, idxmin, n int) int {
	if i == n {
		return idxmin
	} else {
		if T[i] < T[idxmin] {
			idxmin = i
		}
		return min(T, i+1, idxmin, n)
	}
}
