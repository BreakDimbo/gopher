package main

import (
	"fmt"
	"os"
)

// O(n^2)
func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
