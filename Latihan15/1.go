package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func main() {
	var A arrInt
	var i, n int

	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	print(A, n)
}
func print(T arrInt, n int) {
	printarray(T, 0, n)
}

func printarray(T arrInt, i, n int) {
	if i < n {
		fmt.Printf("%v ", T[i])
		printarray(T, i+1, n)
	}
}
