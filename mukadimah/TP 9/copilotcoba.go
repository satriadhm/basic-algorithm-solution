package main

import "fmt"

//quick sort algorithm
func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}
func main() {
	var arr []int
	arr = []int{5, 2, 4, 6, 1, 3}
	quickSort(arr, 0, len(arr)-1)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}
func partition(arr []int, low int, high int) int {
	i := low - 1
	pivot := arr[high]
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

//insertion sort algorithm

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

//maximum element in array

func max(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

//binary search algorithm
func binarySearch(arr []int, l int, r int, x int) int {
	if r >= l {
		mid := l + (r-l)/2
		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		}
		return binarySearch(arr, mid+1, r, x)
	}
	return -1
}
