package main

import "fmt"

func main() {
	var n int
	var a, b int
	sum := 0
	i := 1

	fmt.Scan(&n)
	for i <= n {
		fmt.Scan(&a, &b)
		j := 2
		sum = a
		for j <= b {
			a = a * sum
			j++
		}
		i++
		fmt.Print(a)
	}

}
