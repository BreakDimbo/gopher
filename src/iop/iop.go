package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(content)

	count := countWord(text)
	fmt.Printf("There are %d words in your text. \n", count)
}

func countWord(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
