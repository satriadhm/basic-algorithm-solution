package main

import "fmt"

type Gol struct {
	nama1, nama2 string
	goal, assist int
}
type arrInt [1001]Gol

func main() {
	var arr arrInt
	var gol Gol
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&gol.nama1, &gol.nama2, &gol.goal, &gol.assist)
		arr[i] = gol
	}
	sort(&arr, n)

}
func sort(arr *arrInt, n int) {
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].goal*1000+arr[j].assist > arr[minIdx].goal*1000+arr[minIdx].assist {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
		fmt.Print(arr[i].nama1, arr[i].nama2, " ", arr[i].goal, arr[i].assist, "\n")

	}
}
