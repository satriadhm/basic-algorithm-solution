package main

import "fmt"

func main() {
	var ani1, ani2, ani3 string
	var status bool
	fmt.Scan(&ani1, &ani2, &ani3)
	status = false
	if ani1 == ani2 || ani2 == ani3 || ani3 == ani1 {
		status = true
	}
	fmt.Println(status)
}
