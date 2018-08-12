package main

import (
	"fmt"
	"os"
	"strings"
)

// Echo prints its command-line arguments.
func main() {

	/*
		// 1st version
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = ""
		}
	*/

	/*
		// 2nd version
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
	*/

	// 3rd version
	// fmt.Println(strings.Join(os.Args[1:], " "))

	// ex1.2
	for index, arg := range os.Args {
		fmt.Printf("index: %d, arg: %s\n", index, arg)
	}
}

/*
ex1.3
*/

func echo2(args []string) {
	// 2nd version
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	fmt.Println(strings.Join(args, " "))
}
