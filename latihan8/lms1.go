package main

import "fmt"

func main() {
	var hasil bool
	var x, y, z int
	x = 5
	y = 3
	z = 2
	cekSelisih(x, y, z, &hasil)
	fmt.Println(hasil)
}

func cekSelisih(x, y, z int, hasil *bool) {
	if (x-y == z) || (y-x == z) {
		*hasil = true
	} else {
		*hasil = false
	}

}

func konsekutif(x int) int {
	res := 0
	status := true
	d1 := x % 10
	d0 := d1
	x = x / 10
	for x > 0 && status {
		d1 = x % 10
		status = (d1-d0) == 1 || (d1-d0) == -1
		d0 = d1
		x = x / 10
		if status == true {
			res = x
		} else if status == false {
			res = -1
		}
	}
	return res
}
