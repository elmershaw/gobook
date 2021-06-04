package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseBytes(b []byte) []byte {
	l, r := 0, len(b)
	for l < r-1 {
		left, leftSize := utf8.DecodeRune(b[l:])
		right, rightSize := utf8.DecodeLastRune(b[:r])
		copy(b[l+rightSize:], b[l+leftSize:r-rightSize])
		utf8.EncodeRune(b[l:], right)
		utf8.EncodeRune(b[r-leftSize:], left)
		l += rightSize
		r -= leftSize
	}
	return b
}

func main() {
	b := []byte("哈哈  哈 哈哈  a")
	fmt.Printf("%s\n", b)
	reverseBytes(b)
	fmt.Printf("%s\n", b)
}
