package main

import "fmt"

func main() {
	var n, x int
	res := 0
	i := 0
	fmt.Scan(&n)
	for i < n {
		fmt.Scan(&x)
		for (x > 9) || (x < 0) {
			fmt.Scan(&x)
			res += 1
		}
		i++
	}

	fmt.Print(res)
}
