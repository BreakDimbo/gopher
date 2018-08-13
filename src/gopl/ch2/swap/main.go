package main

import (
	"fmt"
)

func main() {
	x, y := 1, 5
	fmt.Printf("the numbers to be swapped are %d, %d\n", x, y)

	swap(&x, &y)
	swap2(&x, &y)
	fmt.Printf("swap twice over: %d, %d\n", x, y)
}

func swap(x, y *int) {
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
}

func swap2(x, y *int) {
	*x, *y = *y, *x
}
