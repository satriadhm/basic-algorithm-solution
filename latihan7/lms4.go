package main

import "fmt"

func main() {
	var n int
	var tendangan int
	i := 0
	marcus := 0
	degea := 0
	fmt.Scan(&n)
	for i < n {
		fmt.Scan(&tendangan)
		if tendangan%2 == 0 {
			fmt.Print("Tendangan terlalu ke samping")
			degea += 1
		} else if tendangan%5 == 0 {
			fmt.Print("Tendangan terlalu ke atas")
			degea += 1
		} else if tendangan%2 != 0 && tendangan%5 != 0 {
			fmt.Print("Tendangan tepat sasaran")
			marcus += 1
		}
		i++
	}
	fmt.Print("Skor akhir: ", marcus, "untuk Marcus,", degea, "untuk De Gea")
}
