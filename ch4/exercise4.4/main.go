package main

import "fmt"

func rotate(a []int, n int) []int {
	for i := 0; i < n; i++ {
		a = append(a[1:], a[0])
	}
	return a
}

func main() {
	array := []int{0, 1, 2, 3, 4, 5}
	array = rotate(array, 2)
	fmt.Println(array)
}
