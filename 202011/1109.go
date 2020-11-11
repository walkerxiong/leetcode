package main

import (
	"github.com/gbabyX/leetcode/util"
)

func kClosest(points [][]int, K int) [][]int {
	var (
		res [][]int
		m   = make(map[int][]int) // distance => seq
		h   = util.NewIntHeap(func(i, j int) bool { return i <= j }, K)
	)
	for seq, arr := range points {
		distance := arr[0]*arr[0] + arr[1]*arr[1]
		m[distance] = append(m[distance], seq)
		h.Push(distance)
	}
	for i := 0; i < K; {
		v := h.Pop()
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
