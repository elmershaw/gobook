package popcount

func PopCount(x int64) {
	var res int
	for x > 0 {
		if x&1 != 0 {
			res++
		}
		x = x >> 1
	}
}
