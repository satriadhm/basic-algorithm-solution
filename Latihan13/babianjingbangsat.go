package main

import "fmt"

type gol struct {
	tim   string
	total int
}

const NMAX int = 100000

type tabInt [NMAX]gol

func main() {
	var array tabInt
	var n int
	inputData(&array, &n)
	sort(&array, n)
	for i := 0; i < n; i++ {
		fmt.Printf("%v %v\n", array[i].tim, array[i].total)
	}

}

func inputData(A *tabInt, n *int) {
	var tim string
	var menang, kalah, draw int
	fmt.Scanln(&tim, &menang, &kalah, &draw)
	for tim != "None" {
		(*A)[*n].tim = tim
		(*A)[*n].total = menang*3 + kalah*0 + draw*1
		*n++
		fmt.Scanln(&tim, &menang, &kalah, &draw)
	}

}
func sort(arr *tabInt, n int) {
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].total > arr[minIdx].total {
				minIdx = j
			}
		}
		temp_name := (*arr)[i].tim
		temp_total := (*arr)[i].total
		(*arr)[i].tim = (*arr)[minIdx].tim
		(*arr)[i].total = (*arr)[minIdx].total
		(*arr)[minIdx].tim = temp_name
		(*arr)[minIdx].total = temp_total
	}
}
