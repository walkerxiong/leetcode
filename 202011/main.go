package main

import "fmt"

func main() {
	points := [][]int{[]int{3, 3}, []int{5, -1}, []int{5, -1}, []int{-2, 4}}
	var K = 2
	res := kClosest(points, K)
	fmt.Println(res)

}
