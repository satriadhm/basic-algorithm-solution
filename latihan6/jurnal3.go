package main

import "fmt"

func main() {
	var n, num, hitung int
	fmt.Scan(&n)
	i := 0
	hitung = 0
	for i < n {
		fmt.Scan(&num)
		if num < 10 && num >= 0 {
			hitung += num
			i++
		}
	}
	fmt.Print(hitung)
}
