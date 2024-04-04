package main

import "fmt"

const NMAX = 1000000

type tabInt [NMAX]int

var p tabInt

func main() {
	var partai int
	var n int
	for partai != -1 {
		fmt.Scan(&partai)
		p[partai]++
		n++
	}
	array(p, n)
}
func array(arr tabInt, n int) {
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[minIdx] {
				minIdx = j
			}

		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
		
	}


func posisi(t tabInt, n int, nama int) int {
	var sum int
	var ketemu bool
	var k int
	ketemu = false
	for !ketemu && k < len(t) {
		if t[k] == nama {
			ketemu = true
			sum++
		}
		k++
	}
	return sum

}
