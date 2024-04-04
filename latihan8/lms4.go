package main

func main() {
	
}
func valid(x,k int) int{
	count := 0
	if x > 0 && x < k{
		count = k
	}else if x < 0 {
		count = 0
	}else if x > k {
		count = x
	}
	return count
}

func rataan(mean *float64)