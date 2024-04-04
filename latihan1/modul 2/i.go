package main

import "fmt"

func main() {
	var n, jumlah, t1, t2, t3, i int
	fmt.Scan(&n)
	jumlah = 0
	for i = 0; i <= n; i += 1 {
		fmt.Scanln(&t1, &t2, &t3)
		if t1+t2+t3 >= 2 {
			jumlah = jumlah + 1
		}
	}
	fmt.Println(jumlah)

}
