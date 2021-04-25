package popcount

func PopCount(x uint64) int {
	var res int
	for x > 0 {
		res++
		x = x & (x - 1)
	}
	return res
}
