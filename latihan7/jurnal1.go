package main

import "fmt"

func main() {
	var b int
	var prima bool
	fmt.Scan(&b)
	fmt.Print("Faktor :")
	i := 1
	faktor := 0
	for i <= b {
		iteration := 0
		for b%i == 0 && iteration == 0 {
			fmt.Print(i, " ")
			faktor++
			iteration++

		}
		i++
	}
	fmt.Println("")
	prima = faktor == 2
	fmt.Print("Prima :", prima)
}
