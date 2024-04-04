package main

import "fmt"

func main() {
	var tem1, tem2, tem3, tem4 float64
	fmt.Scan(&tem1, &tem2, &tem3, &tem4)
	fmt.Println("Temperatur naik?", tem1 < tem2 && tem2 < tem3 && tem3 < tem4)
	fmt.Println("Temperatur turun?", tem1 > tem2 && tem2 > tem3 && tem3 > tem4)
	fmt.Println("Temperatur tetap?", tem1 == tem2 && tem2 == tem3 && tem3 == tem4)
	fmt.Println("Temperatur naik turun?", tem1 < tem2 || tem2 < tem3 || tem3 < tem4 && tem1 > tem2 && tem2 > tem3 && tem3 > tem4)
}
