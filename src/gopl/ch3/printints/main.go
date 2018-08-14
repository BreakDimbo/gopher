package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Print(intsToString([]int{1, 2, 3}))
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	return buf.String()
}
