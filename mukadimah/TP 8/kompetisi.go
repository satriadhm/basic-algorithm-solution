package main

import "fmt"

type tabInt [1000000]int

func main() {
	var input int
	var ar tabInt
	var n int
	fmt.Scan(&input)
	for input != -5313541 {
		if input == 0 {
			array(&ar, n)
			m := median(&ar, n)
			fmt.Println(m)
		} else {
			ar[n] = input
			n++
		}
		fmt.Scan(&input)
	}
}
func array(arr *tabInt, n int) {
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
func median(arr *tabInt, n int) float64 {
	var med float64
	if n%2 == 0 {
		med = (float64(arr[(n/2)-1]) + float64(arr[n/2])) / 2.0
	} else {
		med = float64(arr[n/2])
	}
	return med
}
