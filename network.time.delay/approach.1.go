package main

import (
	"container/heap"
	"math"
)

type MinHeap [][2]int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Pair struct {
	node, dist int
}

func networkDelayTime(times [][]int, n, k int) int {
	adj := make(map[int][]Pair)

	for _, time := range times {
		source, dest, travelTime := time[0], time[1], time[2]
		adj[source] = append(adj[source], Pair{dest, travelTime})
	}

	dist := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		if i != 0 {
			dist[i] = math.MaxInt
		}
	}

	var h MinHeap
	heap.Init(&h)

	heap.Push(&h, [2]int{k, 0})
	dist[k] = 0

	for len(h) > 0 {
		v := heap.Pop(&h).([2]int)
		n := v[0]
		d := v[1]

		if d > dist[n] {
			continue
		}

		for _, en := range adj[n] {
			nei := en.node
			di := en.dist

			newD := di + d

			if newD < dist[nei] {
				heap.Push(&h, [2]int{nei, newD})
				dist[nei] = newD
			}
		}
	}

	minTime := math.MaxInt

	for j := 1; j < n+1; j++ {
		minTime = max(dist[j], minTime)
	}

	if minTime == math.MaxInt {
		return -1
	}
	return minTime
}

func max(a, b int) int {
	if a < b {
		return a
	}
	return b
}
