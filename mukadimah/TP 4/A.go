package main

import "fmt"

func main() {

	maxmin(1, 2, 3, -1)
}

func MaxMin(a, b, c, d int) (int, int) {
	max := 0
	min := 0

	for _, num := range []int{a, b, c, d} {
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	return max, min

}

func maxmin(a, b, c, d int) {
	max := 0
	min := 0

	for _, num := range []int{a, b, c, d} {
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	fmt.Print(max, min)

}
