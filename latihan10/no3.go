package main

import "fmt"

func main() {
	var a, b int
	menit := 0
	fmt.Scan(&a, &b)
	for {
		menit++
		if menit%a == 0 && menit%b == 0 {
			break
		}
	}
	fmt.Print(menit)
}
