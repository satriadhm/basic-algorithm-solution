package main

import "fmt"

const NMAX = 1000

type arrint [NMAX]int

func main() {
	var n int
	var a arrint

	fmt.Print("Banyak data : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(sumArr(a, 0, n))
}

func sumArr(a arrint, i, N int) int {
	if N == 0 {
		return 0
	} else {
		return a[i] + sumArr(a, i+1, N)
	}
}
