package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var f, g, h, i, j string
	fmt.Scan(&a, &b, &c, &d, &e)
	fmt.Scan(&f, &g, &h, &i, &j)
	fmt.Println("Tepat 2 tahun? ", e-d == 2 && d-c == 2 && c-b == 2 && b-a == 2)
	fmt.Println("Satu keluarga ", f == g && g == h && h == i && i == j && j == f)

}
