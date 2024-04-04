package main

import "fmt"

func main() {
	var angka, energi1, energi2, positif, negatif int
	fmt.Scanln(&angka)

	for i := 0; i <= angka; i++ {
		fmt.Scanln(&energi1, &energi2)
		if energi1 < 0 {
			negatif += energi1
		} else {
			positif += energi1
		}
		if energi2 < 0 {
			negatif += energi2
		} else {
			positif += energi2
		}

	}
	fmt.Println("Negative: ", negatif)
	fmt.Println("Positive: ", positif)

}
