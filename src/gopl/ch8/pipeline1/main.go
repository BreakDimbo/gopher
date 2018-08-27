package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squarer := make(chan int)

	// generate naturls
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// squar it
	go func() {
		for {
			x := <-naturals
			squarer <- x * x
		}
	}()

	// print it
	for {
		fmt.Println(<-squarer)
	}
}
