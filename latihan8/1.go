package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scanln(&a, &b, &c, &d)
	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
func findFactorial(n int, result *int) {
	*result = 1
	for n > 0 {
		*result = *result * n
		n = n - 1
	}
}

func permutation(n int, r int) int {
	var result, fn, fnr int
	findFactorial(n, &fn)
	findFactorial(n-r, &fnr)
	result = fn / fnr
	return result
}

func combination(n int, r int) int {
	var result, fn, fr, fnr int
	findFactorial(n, &fn)
	findFactorial(r, &fr)
	findFactorial(n-r, &fnr)
	result = fn / (fr * fnr)
	return result
}
