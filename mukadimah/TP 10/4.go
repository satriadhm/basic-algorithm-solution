package main

import "fmt"

func main() {
	num := 100
	binar(num)
}

func binar(num int) {
	if num > 1 {
		binar(num / 2)
	}
	fmt.Print(num % 2)
}
