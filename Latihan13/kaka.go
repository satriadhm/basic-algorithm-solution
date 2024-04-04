package main

import "fmt"

type nilai struct {
	nama string
	biji float64
}

type arrInt []nilai

func main() {

}
func inputData(arr *arrInt, n *int) {
	var nama string
	var uts, uas, praktikum float64
	fmt.Scan(&nama, &uts, &uas, &praktikum)
	for (uts < 10 && uts > 0) && (uas < 10 && uas > 0) && (praktikum < 10 && praktikum > 0) {
		(*arr)[*n].nama = nama
		(*arr)[*n].biji = uts*0.4 + uas*0.35 + praktikum*0.25
		*n++
		fmt.Scan(&nama, &uts, &uas, &praktikum)
	}
}
func sort(arr *arrInt, n int) {
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if (*arr)[j].biji > (*arr)[minIdx].biji {
				minIdx = j
			}
		}
		temp_name := (*arr)[i].nama
		temp_total := (*arr)[i].biji
		(*arr)[i].nama = (*arr)[minIdx].nama
		(*arr)[i].biji = (*arr)[minIdx].biji
		(*arr)[minIdx].nama = temp_name
		(*arr)[minIdx].biji = temp_total

	}
}
