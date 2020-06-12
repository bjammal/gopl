package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	counts := make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countElements: %v\n", err)
		os.Exit(1)
	}
	elementCount(counts, doc)
	for k, v := range counts {
		fmt.Println(k, ":", v)
	}
}

func elementCount(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	if n.FirstChild != nil {
		elementCount(counts, n.FirstChild)
	}
	if n.NextSibling != nil {
		elementCount(counts, n.NextSibling)
	}
}
