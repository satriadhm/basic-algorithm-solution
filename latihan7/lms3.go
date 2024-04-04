package main

import "fmt"

func main() {
	var input string
	sum := 0
	for input != "Q" {
		fmt.Scan(&input)
		if input == "W" {
			sum += 3
		} else if input == "D" {
			sum += 1
		} else if input == "L" {
			sum += 0
		}
		fmt.Scan(&input)
	}
	fmt.Print(sum)
}
