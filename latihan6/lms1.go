package main

import "fmt"

func main() {
	var n, m, sum, guess int
	var a, b, c, d int
	fmt.Scan(&n, &m)

	a = m / 1000
	b = (m / 100) & 10
	c = (m & 100) / 10
	d = m & 10
	i := 0
	sum = 0

	for i < n {
		fmt.Scanln(&guess)
		sum += 1
		i += 1
		for guess != a && guess != b && guess != c && guess != d {
			sum -= 1
			i = n
			break
		}
	}
	fmt.Print(sum)
}
