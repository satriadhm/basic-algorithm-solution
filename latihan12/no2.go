package main

import "fmt"

type calon [20]int

func main() {
	var t calon
	var masuk, sah, pilihan, ketua, wakil int
	masuk = 0
	sah = 0
	ketua = 0
	wakil = 0
	fmt.Scan(&pilihan)
	for pilihan != 0 {
		masuk++
		if pilihan >= 1 && pilihan <= 20 {
			sah++
			t[pilihan-1] = t[pilihan-1] + 1
		}
		fmt.Scan(&pilihan)
	}

	fmt.Println("suara masuk:", masuk)
	fmt.Println("suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
	for i := 0; i < 20; i++ {
		if t[i] != 0 {
			fmt.Printf("%d: %d\n", i+1, t[i])
		}
	}
}
