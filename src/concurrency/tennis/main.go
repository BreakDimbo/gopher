package main

import (
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)
	wg.Add(2)

	go player("King", court)
	go player("Queue", court)

	// Start the game
	court <- 1

	// Wait for the game over
	wg.Wait()
}
