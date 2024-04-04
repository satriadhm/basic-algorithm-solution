package main

import "fmt"

func main() {
	var n int
	var tipe bool
	fmt.Scan(&n, &tipe)
	if n <= 10 {
		fmt.Print(0)
	} else if tipe {
		if float32(n)/60 != float32(n/60) {
			fmt.Print(((n / 60) + 1) * 4000)
		} else {
			fmt.Print(n / 60 * 4000)
		}
	} else {
		if n < 60 {
			fmt.Print(20000)
		} else if n < 300 {
			fmt.Print(25000)
		} else {
			fmt.Print(35000)
		}
	}
}
