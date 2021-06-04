package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		words[in.Text()]++
	}
	for word, count := range words {
		fmt.Printf("Word: %s, Count: %d\n", word, count)
	}
}
