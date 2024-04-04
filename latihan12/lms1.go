package main

import "fmt"

type set [999]int

func main() {
	var num, n, x, y, i, ketemu, sum int
	var arr set
	n = 0
	fmt.Scan(&num)
	for n < 999 && num >= 0 {
		arr[n] = num
		n++
		fmt.Scan(&num)
	}

	fmt.Scan(&x, &y)
	fmt.Scan(&ketemu)
	for i = x; i <= y; i++ {
		if arr[i] == ketemu {
			sum++
		}
	}
	fmt.Println(sum)
}
