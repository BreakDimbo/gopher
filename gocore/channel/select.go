package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	rIndex := rand.Intn(10)
	fmt.Println(rIndex)
	intChannels[rIndex] <- 1
	select {
	case _, ok := <-intChannels[0]:
		if !ok {
			intChannels[0] = nil
		}
		fmt.Println("get value from channel 0")
	case <-intChannels[1]:
		fmt.Println("get value from channel 1")
	case <-intChannels[2]:
		fmt.Println("get value from channel 2")
	default:
		fmt.Println("get value from somewhere")
	}
	a := sum
	a()
}

func sum() {
	http.HandleFunc()
}
