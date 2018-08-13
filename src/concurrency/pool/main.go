package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines = 25
	maxPoolSize   = 2
)

type dbConnection struct {
	ID int32
}

func (db *dbConnection) Close() error {
	log.Println("Close: Connection ", db.ID)
	return nil
}

var idCounter int32

func createConnction() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection ", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	pool, err := New(createConnction, maxPoolSize)
	if err != nil {
		log.Println(err)
	}

	for query := 1; query <= maxGoroutines; query++ {
		go func(q int) {
			performQuery(q, pool)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Program", "Shutting Down.")
	pool.Close()
}

func performQuery(query int, pool *Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer pool.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d], CID[%d]\n", query, conn.(*dbConnection).ID)
}
