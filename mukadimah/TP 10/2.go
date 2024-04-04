package main

import "fmt"

func main() { // Nyoba.go
	var array = []int{47, 100, 88, 92, 34, 1, 9}
	reccursive(array)
}

func reccursive(arr []int) {
	n := len(arr)
	if n > 0 {
		reccursive(arr[:n-1])
		fmt.Println(arr[n-1])
	}
}
