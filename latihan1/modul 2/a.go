package main

import "fmt"

func main() {
	var s string
	var a, b int
	fmt.Scan(&s)
	fmt.Scan(&a, &b)
	fmt.Println("Kata = ", s)
	fmt.Println("Jumlah = ", a+b)

}
