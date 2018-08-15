package main

import (
	"fmt"
)

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(removeWithOrder(s, 2))
}

func removeWithOrder(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeWithoutOrder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
