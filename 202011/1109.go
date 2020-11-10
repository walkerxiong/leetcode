package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, K int) [][]int {
	var (
		res [][]int
		m   = make(map[int][]int) // distance => seq
		h   = &IntHeap{}
	)
	heap.Init(h)
	for seq, arr := range points {
		distance := arr[0]*arr[0] + arr[1]*arr[1]
		m[distance] = append(m[distance], seq)
		heap.Push(h, distance)
	}
	fmt.Println(m)
	for i := 0; i < K; {
		v := heap.Pop(h).(int)
		val := m[v]
		if val != nil {
			for _, s := range val {
				if i >= K {
					break
				}
				res = append(res, points[s])
				i++
			}
		}

	}

	return res
}
