package main

import "fmt"

func main() {
	var w, x, y, z float64
	fmt.Scan(&w)
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&z)
	fmt.Println(w/1000.0, w/453.592, w/28.3495)
	fmt.Println(x/1000.0, x/453.592, x/28.3495)
	fmt.Println(y/1000.0, y/453.592, y/28.3495)
	fmt.Println(z/1000.0, z/453.592, z/28.3495)
}
