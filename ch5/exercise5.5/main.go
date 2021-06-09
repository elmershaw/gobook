package exercise55

import (
	"bufio"
	"strings"

	"golang.org/x/net/html"
)

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		words = wordCount(n.Data)
	}
	wt, it := countWordsAndImages(n.FirstChild)
	words, images = wt+words, it+images
	wt, it = countWordsAndImages(n.NextSibling)
	words, images = wt+words, it+images
	return
}

func wordCount(s string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n
}
