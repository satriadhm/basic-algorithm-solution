package main

import "fmt"

func main() {
	var name, nameMax, nameMin string
	var factor, meanMin, meanMax float64

	meanMin = 100
	meanMax = 0
	fmt.Scan(&name)
	for name != "STOP" {
		var p int
		var jumlah, rata2 float64
		fmt.Scan(&p)
		i := 0
		for i < 3*p {
			fmt.Scan(&factor)
			jumlah += factor
			i++
		}
		rata2 = jumlah / (3 * float64(p))
		if rata2 > meanMax {
			meanMax = rata2
			nameMax = name
		}
		if rata2 < meanMin {
			meanMin = rata2
			nameMin = name
		}
		fmt.Scan(&name)
	}
	fmt.Printf("%s %.2f\n", nameMax, meanMax)
	fmt.Printf("%s %.2f", nameMin, meanMin)
}
