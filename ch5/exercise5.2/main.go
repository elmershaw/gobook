package main

import "golang.org/x/net/html"

func nameCounter(counter map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return counter
	}
	if n.Type == html.ElementNode {
		counter[n.Data]++
	}
	counter = nameCounter(counter, n.FirstChild)
	counter = nameCounter(counter, n.NextSibling)
	return counter
}
