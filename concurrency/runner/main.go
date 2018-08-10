package main

import (
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Start working.")

	runner := New(timeout)

	runner.Add(createTask(), createTask(), createTask())

	if err := runner.Start(); err != nil {
		switch err {
		case ErrInterrupt:
			log.Println("Terminate due to interrupt.")
			os.Exit(2)
		case ErrTimeout:
			log.Println("Terminate due to timeout.")
			os.Exit(1)
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
