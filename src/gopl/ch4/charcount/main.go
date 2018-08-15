package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	lcounts := make(map[rune]int)
	dcounts := make(map[rune]int)

	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			lcounts[r]++
		}

		if unicode.IsDigit(r) {
			dcounts[r]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for r, c := range counts {
		fmt.Printf("%q\t%d\n", r, c)
	}

	fmt.Printf("letter\tcount\n")
	for r, c := range lcounts {
		fmt.Printf("%q\t%d\n", r, c)
	}

	fmt.Printf("digit\tcount\n")
	for r, c := range dcounts {
		fmt.Printf("%q\t%d\n", r, c)
	}

	fmt.Printf("\nlen\tcount\n")
	for l, c := range utflen {
		fmt.Printf("%d\t%d\n", l, c)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
