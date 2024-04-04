package main

import "fmt"

func main() {
	var berat, kg, gram, maincost, detailcost int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)
	kg = berat / 1000
	gram = berat % 1000
	maincost = kg * 10000
	if kg > 10 {
		detailcost = 0
	} else if gram >= 500 {
		detailcost = gram * 5
	} else {
		detailcost = gram * 15
	}
	fmt.Println("Detail berat: ", kg, "kg +", gram, "gr")
	fmt.Println("Detail biaya: Rp.", maincost, "+ Rp.", detailcost)
	fmt.Println("Total biaya: Rp.", maincost+detailcost)

}
