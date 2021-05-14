package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func getComma(s string) string {
	strs := strings.Split(s, ".")
	str := strs[0]
	if str[0] == '+' || str[0] == '-' {
		return fmt.Sprintf("%c"+string(str[1:]), str[0]) + "." + strs[1]
	}
	return comma(str) + "." + strs[1]
}

func main() {
	fmt.Println(getComma("-1234.25"))
}
