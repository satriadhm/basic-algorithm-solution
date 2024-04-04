package main

type tabInt [100000]int

func CariSekuensial(t tabInt, n int) int {
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
