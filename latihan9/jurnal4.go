package main

import "fmt"

func main() {
	var walk bool
	var i, input, before, day, sum int

	walk = true
	i = 0
	day = 0

	for walk {
		day += 1
		fmt.Print("Hari ", day, " : ")
		fmt.Scan(&input)
		for (input >= 0 && input <= 100) == false {
			fmt.Print("Hari ", day, " : ")
			fmt.Scan(&input)
		}
		sum += input
		if input <= before {
			i += 1
		}
		if i == 3 {
			walk = false
		}
		before = input
	}
	mean := float64(sum) / float64(day)
	fmt.Println("Museum buka selama", day, "hari")
	fmt.Printf("Rata-rata pengunjung %.2f", mean)
}
