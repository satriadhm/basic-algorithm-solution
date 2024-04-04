package main

import "fmt"

type employee struct {
	name    string
	bonus   float64
	year    int
	married bool
}

func isidata(P *employee) {
	fmt.Scan(&P.name, &P.bonus, &P.year, &P.married)
}

func hitungBonus(P *employee) {
	var temp1, temp2 float64
	temp1 = 2022 - (float64((*P).year))
	if temp1 >= 10 {
		temp1 *= 10000
	} else {
		temp1 = 0
	}
	if (*P).married {
		temp2 = 1.5 * temp1
	}
	(*P).bonus = temp1 + temp2
}
