package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func visit(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		for _, a := range n.Attr {
			fmt.Println(a.Val)
		}
	}
	visit(n.FirstChild)
	visit(n.NextSibling)
}
