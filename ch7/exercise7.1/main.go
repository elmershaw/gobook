package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	s := "Hello, World!\nHello, 世界！"

	var wc WordCounter
	fmt.Fprint(&wc, s)
	fmt.Println(wc)

	var lc LineCounter
	fmt.Fprint(&lc, s)
	fmt.Println(lc)
}

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*wc++
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}
