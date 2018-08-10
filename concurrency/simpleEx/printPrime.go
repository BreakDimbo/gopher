package main

import (
	"fmt"
)

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 10000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Complete ", prefix)
}
