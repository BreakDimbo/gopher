package main

import (
	"fmt"
)

func main() {
	data := []int{4, 2, 5, 6, 77, 8, 9, 23}
	Sort(data)
	fmt.Println(data)
}

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts value in place
func Sort(values []int) {
	var root *tree
	for _, value := range values {
		root = add(root, value)
	}

	appendValue(values[:0], root)
}

func appendValue(values []int, t *tree) []int {
	if t != nil {
		values = appendValue(values, t.left)
		values = append(values, t.value)
		values = appendValue(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = &tree{value: value}
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
