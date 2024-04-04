package main

import "fmt"

const NMAX = 1000

type arrint [NMAX]int

func main() {
	var n int
	var A arrint

	fmt.Print("Input n : ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	fmt.Println("Palindrom = ", palindrom(A, 0, n-1, n))
}

func palindrom(A arrint, i, j, n int) bool {
	if i == n/2 {
		return A[i] == A[j]
	} else {
		return A[i] == A[j] && palindrom(A, i+1, j-1, n)
	}

}
