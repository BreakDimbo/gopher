package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

const (
	numberGoutines = 4
	taskLoad       = 10
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create a buffered channel to manage load tasks
	tasks := make(chan string, taskLoad)

	// Launch goroutine workers to handle the task
	wg.Add(numberGoutines)
	for gr := 1; gr <= numberGoutines; gr++ {
		go worker(tasks, gr)
	}

	// Send task to the channel
	for task := 1; task <= taskLoad; task++ {
		tasks <- fmt.Sprintf("Task : %d", task)
	}

	// Close channel won't affect the receiver
	close(tasks)

	wg.Wait()
}
