package main

type arr [10000]int

func terurutA(A *arr, n int) {
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
