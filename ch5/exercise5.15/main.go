package main

import (
	"fmt"
	"math"
	"os"
)

func max(args ...int) (res int, err error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("empty args")
	}
	res = math.MinInt64
	for _, integer := range args {
		if integer > res {
			res = integer
		}
	}
	return
}

// func maxWithOne(ncsry int, args ...int) (res int) {
// 	res = ncsry
// 	for _, integer := range args {
// 		if integer > res {
// 			res = integer
// 		}
// 	}
// 	return
// }

func min(args ...int) (res int, err error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("empty args")
	}
	res = math.MaxInt64
	for _, integer := range args {
		if integer < res {
			res = integer
		}
	}
	return
}

// func minWithOne(ncsry int, args ...int) (res int) {
// 	res = ncsry
// 	for _, integer := range args {
// 		if integer < res {
// 			res = integer
// 		}
// 	}
// 	return
// }

func main() {
	maxInt, err := max(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(maxInt)
	minInt, err := min()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(minInt)
}
