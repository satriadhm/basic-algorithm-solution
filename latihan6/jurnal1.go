package main

import "fmt"

func main() {
	var i, n, a, r, ar int
	fmt.Scan(&n, &a, &r)
	i = 1
	ar = a * r
	fmt.Print(0)
	for i <= n {
		fmt.Print(" + ", i*ar)
		i++
	}

}
