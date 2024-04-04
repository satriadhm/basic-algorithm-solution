package main

import "fmt"

type tabGol [2002]int

func inputdata(t *tabGol, n *int) {
	var win int
	fmt.Scanf(&win)
	for win >= 0 {
		t[*n] = win

	}
}
