package main

import "fmt"

type set [123]int

func main() {
	var sum, k, num, i, max1, min1, max2, min2, j int
	var arr set
	fmt.Scan(&sum, &k)
	fmt.Scan(&num)
	for i < 123 && i <= sum {
		arr[i] = num
		i++
		fmt.Scan(&num)
	}
	if k > 0 {
		for i = 0; i <= k; i++ {
			if arr[i] < min1 {
				arr[i] = min1
			} else if arr[i] > max1 {
				arr[i] = max1
			}
		}
		for j = k; j < sum; j++ {
			if arr[j] < min2 {
				arr[j] = min2
			} else if arr[j] > max2 {
				arr[j] = max2
			}
			fmt.Print(min1, max1)
			fmt.Print(min2, max2)

		}
	} else {
		fmt.Print("tidak ditemukan")
	}
}
