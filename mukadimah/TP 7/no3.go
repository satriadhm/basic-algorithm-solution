package main

type tabInt [999]int

func ascending(tab tabInt, n, x int) int {
	var left, right, mid int
	var found int
	left = 1
	right = n
	found = -1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x < tab[mid] {
			right = mid - 1
		} else if x > tab[mid] {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found

}

func descending(tab tabInt, n, x int) int {
	var left, right, mid int
	var found int
	left = 1
	right = n
	found = -1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x > tab[mid] {
			right = mid - 1
		} else if x < tab[mid] {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found

}
