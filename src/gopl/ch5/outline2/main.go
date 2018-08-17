package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	s := "123$fooooks$foo"
	fmt.Print(expand(s, f))
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth--
	}
}

func expand(s string, f func(string) string) string {
	i := strings.Index(s, "$foo")
	if i < 0 {
		return s
	}

	if len(s) <= 4 {
		return s[:i] + f("foo")
	}

	s = s[:i] + f("foo") + s[i+4:]
	return expand(s, f)
}

func f(s string) string {
	return "break"
}
