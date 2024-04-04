package main

import "fmt"

func main() {
	var dadu1, dadu2 int
	var sum int
	sum = 0
	for {
		fmt.Scan(&dadu1, &dadu2)
		if dadu1%2 == 1 && dadu2%2 == 1 {
			sum += 1
		}
		if dadu1%2 == 0 && dadu2%2 == 0 {
			break
		}

	}
	fmt.Print(sum)
}
