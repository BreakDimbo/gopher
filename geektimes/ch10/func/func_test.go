package func_test

import (
	"fmt"
	"testing"
	"time"
)

func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Printf("time spent: %fs\n", time.Since(start).Seconds())
		return ret
	}
}

func slowFn(op int) int {
	time.Sleep(2 * time.Second)
	return op
}
func TestTsFn(t *testing.T) {
	tsSF := timeSpent(slowFn)
	t.Log(tsSF(10))

}
