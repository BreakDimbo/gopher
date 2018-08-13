package main

import (
	"fmt"
	"time"
)

func Runner(baton chan int) {
	var newRunner int

	// Waiting for the baton
	runner := <-baton

	fmt.Printf("Runner %d running with Baton.\n", runner)

	// If not the final runner, set the next runner
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the Line.\n", newRunner)
		go Runner(baton)
	}

	// Simulatin Running
	time.Sleep(1 * time.Second)

	// if the final runner, finish the goroutine
	if runner == 4 {
		fmt.Printf("Runner %d finished, Race Over.\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d pass baton to runner %d\n", runner, newRunner)
	baton <- newRunner
}
