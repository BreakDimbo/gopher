package main

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner runs a set of functions within a given timeout and can be
// shutdown on a system operation interrupt signal
type Runner struct {
	// Receive interrupt signal from system operation
	interrupt chan os.Signal

	// Receive complete message or interrupt error
	complete chan error

	// Receive timeout message
	timeout <-chan time.Time

	// The functions which will be executed orderly
	tasks []func(int)
}

// ErrTimeout is returned when a value is received on the timeout
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt is returned when an event from the OS is received
var ErrInterrupt = errors.New("received interrupt")

// New return a Runner with timeout d
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add attachs tasks to the Runner. A task is a function that
// takes an int ID
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all tasks and monitors channel event
func (r *Runner) Start() error {
	// set os signal to our interrupt channel
	signal.Notify(r.interrupt, os.Interrupt)

	// runs all tasks in another goroutine
	go func() {
		r.complete <- r.run()
	}()

	// monitor the channel
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

// run all the tasks
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

// receive interrupt event from os
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		// Stop receive any further signals
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
