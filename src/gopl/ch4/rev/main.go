package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Print(rotate(s, 3))
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, i int) []int {
	if i >= len(s) {
		fmt.Println("i > length of s error")
		return nil
	}

	return append(s[i:], s[:i]...)
}
