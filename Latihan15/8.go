package main

import "fmt"

const NMAX = 1000

type arrint [NMAX]int

func main() {
	var A arrint
	var i, n, x int

	fmt.Scan(&n, &x)
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	fmt.Println(binarysearch(A, n, x))
}

func binarysearch(I arrint, n, x int) int {
	return search(I, 0, n-1, x)
}

func search(I arrint, left, right, x int) int {
	var mid int = (left + right) / 2

	if left > right {
		return -1
	} else if x == I[mid] {
		return mid
	} else {
		if x > I[mid] {
			return search(I, left, mid-1, x)
		} else {
			return search(I, mid+1, right, x)
		}
	}
}
