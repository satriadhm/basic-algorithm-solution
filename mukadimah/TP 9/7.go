package main

type arr [10000]int

func terurutB(a *arr, n int) {
	var pass, i, temp int
	pass = 1
	for pass < n {
		i = pass
		temp = a[pass]
		for i > 0 && temp > a[i-1] {
			a[i] = a[i-1]
			i--
		}
		a[i] = temp
		pass++
	}

}
