package main

import "fmt"

func main2() {
	var n int
	var x, y string
	fmt.Print("N :")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Bunga", i+1, ":")
		fmt.Scan(&x)
		y = y + x
	}
	fmt.Println("Pita :")
	fmt.Println(y)

}
