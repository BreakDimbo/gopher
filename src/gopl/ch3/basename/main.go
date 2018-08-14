package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a/b/c.go"
	s = basename(s)
	fmt.Print(s)
}

// base name remove the directory compoment and a .suffix
func basename(s string) string {
	// Discard every thing before last '/'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything befor last .
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basenameImprove(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
