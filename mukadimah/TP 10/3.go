package main

import "fmt"

var sum = 0

func main() { // Nyoba.go
	var num = 4710395
	rekursif(num)
	fmt.Println(sum)
}

func rekursif(num int) {
	if num > 0 {
		rekursif(num / 10)
		sum++
	}
}
