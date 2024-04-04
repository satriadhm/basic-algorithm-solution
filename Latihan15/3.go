package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Input bilangan decimal : ")
	fmt.Scan(&bilangan)

	fmt.Println("", convertdectobin(bilangan))
}

func convertdectobin(n int) string {
	if n == 0 {
		return "0"
	} else if n == 1 {
		return "1"
	} else if n%2 == 0 {
		return convertdectobin(n/2) + "0"
	} else {
		return convertdectobin(n/2) + "1"
	}
}
