package main

import "fmt"

type tabInt [26]int

func maksimum(T tabInt, n int) int {
	max := T[0]
	for i := 1; i <= n; i++ {
		if max < T[i] {
			max = T[i]
		}
	}
	return max
}
func minimum(T tabInt, n int) int {
	min := T[0]
	for i := 1; i <= n; i++ {
		if min > T[i] {
			min = T[i]
		}
	}
	return min
}
func main() {
	var A tabInt = [26]int{10, 20, 3, 4, 5}
	var n int = 5
	fmt.Println(maksimum(A, n))
	fmt.Println(minimum(A, n))
}
