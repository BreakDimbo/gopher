package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	fmt.Println("Create goroutines.")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to finish.")
	wg.Wait()

	fmt.Println("Terminating program.")
}
