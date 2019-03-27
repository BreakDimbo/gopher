package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	a, b := 1, 1
	t.Log(a)
	for i := 0; i < 7; i++ {
		t.Log(" ", b)
		a, b = b, a+b
	}
	t.Log()
}
