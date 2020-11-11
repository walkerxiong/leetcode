package util

import (
	"fmt"
	"testing"
)

func TestHeap_Push(t *testing.T) {
	h := NewIntHeap(func(i, j int) bool { return i <= j }, 3)
	h.Push(1)
	h.Push(6)
	h.Push(9)
	h.Push(5)
	h.Push(2)
	h.Push(10)
	h.Push(8)
	h.Push(0)
	fmt.Println(h.Data)
}
