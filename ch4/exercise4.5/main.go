package main

import "fmt"

func removeReplica(a []int) []int {
	i := 0
	for i < len(a) {
		if i < len(a)-1 && a[i] == a[i+1] {
			a = remove(a, i)
			continue
		}
		i++
	}
	return a
}

func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	array := []int{0, 1, 1, 1, 1, 5}
	array = removeReplica(array)
	fmt.Println(array)
}
