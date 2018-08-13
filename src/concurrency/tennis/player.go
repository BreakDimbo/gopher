package main

import (
	"fmt"
	"math/rand"
)

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a number randomly
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s missed the ball.\n", name)
			close(court)
			return
		}

		// Display the hit number
		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++

		// Hit the ball back
		court <- ball
	}
}
