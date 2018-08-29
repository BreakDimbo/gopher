package main

import (
	"sync"
)

func main() {

}

var (
	lock    sync.Mutex
	balance int
)

func Deposit(amount int) {
	lock.Lock()
	balance += amount
	lock.Unlock()
}

func Balance() int {
	lock.Lock()
	defer lock.Unlock()
	return balance
}
