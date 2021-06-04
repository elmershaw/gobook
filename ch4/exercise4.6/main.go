package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func replaceSpace(b []byte) []byte {
	for i := 0; i < len(b); {
		first, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(first) {
			second, _ := utf8.DecodeRune(b[i+size:])
			if unicode.IsSpace(second) {
				copy(b[i:], b[i+size:])
				b = b[:len(b)-size]
			}
		}
		i += size
	}
	return b
}

func main() {
	b := []byte("哈哈  哈 哈哈  a")
	fmt.Printf("%s\n", b)
	b = replaceSpace(b)
	fmt.Printf("%s\n", b)
}
