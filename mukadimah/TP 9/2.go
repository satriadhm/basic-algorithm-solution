package main

import "fmt"

type tabInt [100000]int

func tambahData(arr *tabInt, n *int) {
	var input int
	i := 0
	for true {
		fmt.Scan(&input)
		if input == 9999 {
			break
		}
		arr[i] = input
		i++
	}
	*n = i
}
