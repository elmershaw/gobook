package main

import (
	"bytes"
	"fmt"
)

func commaWithBuf(str string) string {
	n := len(str)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if i != 0 && i != n-1 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(str[i])
	}
	return buf.String()
}

func main() {
	fmt.Println(commaWithBuf("1223"))
}
