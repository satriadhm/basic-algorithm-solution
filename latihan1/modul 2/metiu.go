package main

import "fmt"

func main() {
	fmt.Println("=================================================")
	fmt.Println("Hallo semuanya !!!!!")
	fmt.Println("Identitas saya adalah sebagai berikut: ")
	var namadepan string
	var namabelakang string
	fmt.Print("Nama Depan : ")
	fmt.Scanln(&namadepan)
	fmt.Print("Nama Belakang : ")
	fmt.Scanln(&namabelakang)
	fmt.Println("Namaku adalah : ", namadepan, namabelakang)

}
