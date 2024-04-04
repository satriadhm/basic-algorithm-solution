package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	fmt.Println(count_digit(bilangan))
}

func count_digit(bilangan int) int {
	if bilangan == 0 {
		return 0
	} else {
		return 1 + count_digit(bilangan/10)
	}
}
