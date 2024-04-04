package main

import "fmt"

func main() {
	var N int
	a := ""
	fmt.Scan(&N)
	for {
		if N%2 == 1 {
			a = "1" + a
		} else if N%2 == 0 {
			a = "0" + a

		}
		N = N / 2
		if N <= 0 {
			break
		}

	}
	fmt.Println(a)
}
