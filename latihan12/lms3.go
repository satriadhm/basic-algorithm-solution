package main

import "fmt"

type set [123]int

func main() {
	var sum, find, num, max, count int
	var arr set
	fmt.Scan(&sum, &find)
	for i := 0; i < sum; i++ {
		fmt.Scan(&num)
		if num > max && num%2 == 0 {
			max = num
		}
		arr[i] = num

	}
	for i := 0; i < sum; i++ {
		if arr[i] == find {
			count++
		}
	}
	if count >= 1 {
		fmt.Print(max)
	} else {
		fmt.Print("tidak ditemukan")
	}

}
