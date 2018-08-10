package main

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool provide the resources
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrChanClosed = errors.New("channel has been closed.")

func New(fn func() (io.Closer, error), size int) (*Pool, error) {
	if size < 1 {
		return nil, errors.New("channel size is less than 1.")
	}

	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   fn,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire: Shanred Resource.")
		if !ok {
			return nil, ErrChanClosed
		}
		return r, nil
	default:
		log.Println("Acquire New Resource.")
		return p.factory()
	}
}

func (p *Pool) Close() {
	// Secure the operation with the Release operation
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		// If the pool has been closed, do noting
		return
	}

	p.closed = true

	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue.")
	default:
		log.Println("Release", "Closing.")
		r.Close()
	}
}
