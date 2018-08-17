package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	outline(nil, doc)
	for elem, count := range countElements(nil, doc) {
		fmt.Printf("element: %s\tcount: %d\n", elem, count)
	}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func countElements(counts map[string]int, n *html.Node) map[string]int {
	if counts == nil {
		counts = make(map[string]int)
	}

	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countElements(counts, c)
	}
	return counts
}

