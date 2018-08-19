package main

import (
	"fmt"
	"gopl/ch5/links"
	"log"
	"os"
)

func main() {
	breathFirst(crawl, os.Args[1:])
}

func breathFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil

		for _, url := range items {
			if !seen[url] {
				seen[url] = true
				worklist = append(worklist, f(url)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
