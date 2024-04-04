package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	check := (((n % 10) - (n / 10 % 10)) == 2) || (((n % 10) - (n / 10 % 10)) == -2)
	for (n < 9) && check {
		check = (((n % 10) - (n / 10 % 10)) == 2) || (((n % 10) - (n / 10 % 10)) == -2)
		n = n / 10
	}
	fmt.Print(check)
}
