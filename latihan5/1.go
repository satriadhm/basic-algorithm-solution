package main

import "fmt"

func main() {
	var xA1, xA2, yA1, yA2 int
	var xB1, xB2, yB1, yB2 int
	var xC1, xC2, yC1, yC2 int
	fmt.Scan(&xA1, &yA1, &xA2, &yA2)
	fmt.Scan(&xB1, &yB1, &xB2, &yB2)
	fmt.Scan(&xC1, &yC1, &xC2, &yC2)
	fmt.Println(float64(yA1-yA2) / float64(xA1-xA2))
	fmt.Println(float64(yB1-yB2) / float64(xB1-xB2))
	fmt.Println(float64(yC1-yC2) / float64(xC1-xC2))
}
