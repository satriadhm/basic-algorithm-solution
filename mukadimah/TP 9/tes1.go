package main

import "fmt"

type tabInt [9]int

func main() {
	var T tabInt
	var n int
	tambahData(&T, &n)
	// fmt.Scan(&vv)
	insertion(&T, n)
	fmt.Print(T)
}
func tambahData(arr *tabInt, n *int) {
	var input int
	i := 0
	for true {
		fmt.Scan(&input)
		if input == 9999 {
			break
		}
		arr[i] = input
		i++
	}
	*n = i
}
func sequential(t tabInt, n int) int {
	i := 0
	for i := 0; i < len(t); i++ {
		if t[i] == n {
			break
		}
		i++
	}
	if i == len(t) {
		return -1
	} else {
		return i
	}
}
func minimum(tab tabInt, n int) int {
	var min int
	var k int
	var idx int
	min = tab[0]
	k = 1
	for k < n {
		if min > tab[k] {
			min = tab[k]
			idx = k
		}
		k++
	}
	return idx

}
func selection(A *tabInt, n int) {
	var pass, idx, i, temp int
	pass = 0
	for pass < n-1 {
		idx = pass
		i = idx + 1
		for i < n {
			if (*A)[i] < (*A)[idx] {
				idx = i
			}
			i++
		}
		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp
		pass = pass + 1
	}
}
func insertion(a *tabInt, n int) {
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
