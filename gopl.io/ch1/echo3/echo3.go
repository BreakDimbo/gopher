package main

import (
	"fmt"
	"os"
	"strings"
)

// O(n)
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
