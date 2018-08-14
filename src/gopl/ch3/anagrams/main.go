package main

import (
	"fmt"
)

func main() {
	a := "apple"
	b := "elppa"
	fmt.Print(isAnagrams(a, b))
}

func isAnagrams(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] != b[n-1-i] {
			return false
		}
	}
	return true
}
