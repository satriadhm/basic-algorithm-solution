package main

import "fmt"

func main() {
	var jum int
	i := 0
	total := 0
	for i < 6 {
		fmt.Scan(&jum)
		if jum > 1000 {
			jum = 1000
			total += jum
		} else if jum < 1000 && jum > 0 {
			total += jum
		} else if jum > 1000 && jum < 0 {
			jum = 0
		}
		i++
	}
	fmt.Print(total)
}
