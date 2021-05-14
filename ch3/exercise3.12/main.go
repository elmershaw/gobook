package main

import (
	"fmt"
	"sort"
	"strings"
)

type unsoredBytes []byte

func (b unsoredBytes) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b unsoredBytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b unsoredBytes) Len() int {
	return len(b)
}

func isMutualShuffled(s1 string, s2 string) bool {
	bytes1, bytes2 := []byte(s1), []byte(s2)
	sort.Sort(unsoredBytes(bytes1))
	sort.Sort(unsoredBytes(bytes2))
	return strings.Compare(string(bytes1), string(bytes2)) == 0
}

func main() {
	fmt.Print(isMutualShuffled("214", "321"))
}
