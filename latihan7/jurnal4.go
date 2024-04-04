package main

import "fmt"

func main() {
	var member string
	var poin1, poin2, poin3, poin4, poin5 int
	var diskon, cashback float64
	fmt.Print(&member, &poin1, &poin2, &poin3, &poin4, &poin5)
	if poin1%2 != 0 && poin2%2 != 0 && poin3%2 != 0 && poin4%2 != 0 && poin5%2 != 0 {
		diskon = 1.7 * float64(poin3+poin4+poin5)
		cashback = 0
	} else if poin1%2 == 0 && poin2%2 == 0 && poin3%2 == 0 && poin4%2 == 0 && poin5%2 == 0 {
		diskon = 0
		cashback = 3.1 * float64(poin1+poin2+poin3)
	} else {
		diskon = 1.7 * float64(poin3+poin4+poin5)
		cashback = 3.1 * float64(poin1+poin2+poin3)
	}
	if member == "yes" {
		diskon += diskon * 0.15
		cashback += cashback * 0.15
	}
	if cashback > 35 {

		cashback = 35
	}
	if diskon > 50 {
		diskon = 50
	}
	fmt.Println(cashback, diskon)
}
