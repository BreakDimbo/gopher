package constant_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	a := 7 // 0111
	// 按位清零
	a = a &^ Writable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
 