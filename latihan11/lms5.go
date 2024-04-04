package main

import "fmt"

type tabInt [15]int

func carinilaitengah(T tabInt, n int) {
	index := n / 2
	var even float64 = float64(T[index]-T[index-1]) / 2
	if n%2 == 0 {
		fmt.Print(T[index])
	} else {
		fmt.Print(even)
	}

}
