package main

import "fmt"

func main() {
	var x, n int
	fmt.Scan(&x, &n)
	kelipatan(x, n)
}

//rekursif kelipatan
func kelipatan(x, n int) {
	lipat(x, x, n)

}

func lipat(x, xn, n int) {
	if xn < n {
		fmt.Print(xn)
		lipat(x, xn+x, n)
	}
}
