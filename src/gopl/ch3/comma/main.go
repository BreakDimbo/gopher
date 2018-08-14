package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "122234.5343"
	fmt.Print(commaEx311(s, true))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaEx310(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])
		if (n-1-i)%3 == 0 && (n-1-i) >= 3 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func commaEx311(s string, isFloat bool) string {
	var digit string
	if isFloat {
		point := strings.LastIndex(s, ".")
		digit = s[point:]
		s = s[:point]
	}

	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])
		if (n-1-i)%3 == 0 && (n-1-i) >= 3 {
			buf.WriteByte(',')
		}
	}
	return buf.String() + digit
}
