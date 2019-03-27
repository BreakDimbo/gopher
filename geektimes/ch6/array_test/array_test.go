package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [...]int{1, 2, 3, 4}
	t.Log(arr[1], arr[2])
	t.Log(arr1)
}

func TestArrayTravel(t *testing.T) {
	arr1 := [...]int{1, 2, 3, 4}
	for _, v := range arr1 {
		t.Log(v)
	}
}
