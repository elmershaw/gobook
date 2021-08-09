package main

import "golang.org/x/net/html"

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var res []*html.Node
	if doc.Type == html.ElementNode {
		for _, tag := range name {
			if doc.Data == tag {
				res = append(res, doc)
			}
		}
	}
	if doc.FirstChild != nil {
		res = append(res, ElementsByTagName(doc.FirstChild, name...)...)
	}
	if doc.NextSibling != nil {
		res = append(res, ElementsByTagName(doc.NextSibling, name...)...)
	}
	return res
}
