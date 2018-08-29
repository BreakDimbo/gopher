// package memo provide a concurrency-unsafe
// memorization of a function of type Func.
package memo4

import (
	"sync"
)

// A Memo caches the results of calling a Func
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type entry struct {
	result
	ready chan struct{}
}

// Func is the type of the function to memorize
type Func func(string) (interface{}, error)

// result is the result of calling the Func
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

//
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.value, e.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()

		<-e.ready
	}

	return e.value, e.err
}
