// Dupex1.4 prints the count and text of lines that appear more than once
// in the input. It pr int t he names of al l files in w hich e ach duplic ate d line o cc urs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineInfo struct {
	count     int
	filenames map[string]int
}

func main() {
	counts := make(map[string]*lineInfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 err %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n.count > 1 {
			var filenames string
			var sep string
			for filename := range n.filenames {
				filenames += sep + filename
				sep = ","
			}
			fmt.Printf("%s\t%d\t%s\n", filenames, n.count, line)
		}
	}
}

func countLines(f *os.File, counts map[string]*lineInfo) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = &lineInfo{filenames: make(map[string]int)}
		}
		counts[input.Text()].count++
		counts[input.Text()].filenames[f.Name()] = 1
	}
}
