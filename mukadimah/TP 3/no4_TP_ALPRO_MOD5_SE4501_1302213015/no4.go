package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a, b int
	var c, d int
	fmt.Scan(&a, &b)
	rand.Seed(int64(a))
	c = rand.Intn(4)
	d = rand.Intn(4)
	fmt.Println(c, d)
	if c == d {
		fmt.Print("Pemenang adalah pak pras")
	} else if b == d && c == d {
		fmt.Println("Seri")
	} else if b == d {
		fmt.Println("Pemenang adalah pemain")
	} else {
		fmt.Println("Tidak ada pemenang")
	}
}
