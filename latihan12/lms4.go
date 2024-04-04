package main

import (
	"fmt"
)

type set [999]int

func maximum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minimum(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	var num, sum, find, min, max, min_idx, max_idx, i int
	var arr set
	var cond string = "tidak ditemukan"
	fmt.Scan(&sum, &find)
	for i = 0; i < sum; i++ {
		fmt.Scan(&num)
		arr[i] = num
	}
	if num == 0 {
		min = num
	}
	if num < min {
		min = num
		min_idx = i
	} else if num > max {
		max = num
		min_idx = i
	}
	if find%2 == 1 {
		fmt.Println("Bilangan ganjil")

	} else {
		for i = minimum(min_idx, max_idx); i <= maximum(min_idx, max_idx) && cond == "tidak ditemukan"; i++ {
			if arr[i] == find {
				cond = "ditemukan"
			}
		}
		fmt.Println(cond)
	}
}
