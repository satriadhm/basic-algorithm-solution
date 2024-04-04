package main

import (
	"fmt"
)

func main1() {
	var benar int
	var a, b, c, d string
	benar = 0
	for i := 0; i < 5; i++ {
		fmt.Print("Percobaan", i+1, ":")
		fmt.Scan(&a, &b, &c, &d)
		if a == "merah" && b == "kuning" && c == "hijau" && d == "ungu" {
			benar++
		}
	}
	if benar == 5 {
		fmt.Println("BERHASIL:true")
	} else {
		fmt.Println("BERHASIL :false")
	}
}
