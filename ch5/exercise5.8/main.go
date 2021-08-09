package main

import "golang.org/x/net/html"

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, findElement, findElement)
}

func forEachNode(n *html.Node, id string, pre, post func(*html.Node, string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node := forEachNode(c, id, pre, post); node != nil {
			return node
		}
	}

	if post != nil {
		if post(n, id) {
			return n
		}
	}

	return nil
}

func findElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	return false
}
