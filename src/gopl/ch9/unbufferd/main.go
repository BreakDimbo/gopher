package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan rune)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("start send %d to ch\n", i)
			ch <- i
			fmt.Printf("end send %d to ch\n", i)
		}
	}()

	go func() {
		for _, c := range "helloworld" {
			fmt.Printf("start send %q to ch2\n", c)
			ch2 <- c
			fmt.Printf("end send %q to ch2\n", c)
		}
	}()

	for {
		select {
		case i := <-ch:
			time.Sleep(1 * time.Second)
			fmt.Printf("receive %d over\n", i)
		case r := <-ch2:
			fmt.Printf("receive %q over\n", r)
		}
	}
}
