package main

import "fmt"

func main() {
	var taun1, taun2, taun3, taun4 int
	var status bool

	fmt.Scan(&taun1, &taun2, &taun3, &taun4)
	status = ((taun1%400 == 0 || taun1%4 == 0) && taun1&100 != 0) && ((taun2%400 == 0 || taun2%4 == 0) && taun2&100 != 0) && ((taun3%400 == 0 || taun3%4 == 0) && taun3&100 != 0) && ((taun4%400 == 0 || taun4%4 == 0) && taun4&100 != 0)
	fmt.Println(status)
}
