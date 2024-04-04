package main

import "fmt"

func main() {
	var n int
	leng1, leng2, leng3, leng4 := true, true, true, true
	i := 0
	fmt.Scan(&n)
	for (i < n) && (leng1 && leng2 && leng3 && leng4) {
		fmt.Scan(&leng1, &leng2, &leng3, &leng4)
		i++
	}
	fmt.Print(i == n)
}
