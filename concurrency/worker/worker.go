package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			// Tasks has been done
			fmt.Printf("Worker %d, shutting donw!\n", worker)
			return
		}

		fmt.Printf("Worker %d, start processing task %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker %d, Complete %s\n", worker, task)
	}
}
