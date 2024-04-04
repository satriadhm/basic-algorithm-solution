package main

import "fmt"

func main() {
	var i, n, g int
	var m, s, t string
	fmt.Scan(&n, &g, &s, &t)
	m = ""
	i = 1
	for i <= n {
		if g%2 == 0 {
			m = m + s
		} else {
			m = m + t
		}
		g++
		i++
	}
	fmt.Print(m)
}
