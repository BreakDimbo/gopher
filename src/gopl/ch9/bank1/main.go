package main

import (
	"fmt"
)

func main() {
	Deposit(100)
	fmt.Printf("Balance: %d\n", Balance())
	fmt.Printf("WithDraw: %t\n", WithDraw(500))
	fmt.Printf("Balance: %d\n", Balance())
}

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan *withDraw)

type withDraw struct {
	amount int
	result chan bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func WithDraw(amount int) bool {
	w := &withDraw{
		amount: amount,
		result: make(chan bool),
	}
	withdraws <- w
	return <-w.result
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw := <-withdraws:
			balance -= withdraw.amount
			withdraw.result <- balance >= 0
		}
	}
}

func init() {
	go teller()
}
