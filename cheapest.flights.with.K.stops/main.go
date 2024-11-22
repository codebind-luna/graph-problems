package main

import (
	"container/heap"
	"math"
)

type pair struct {
	dst   int
	cost  int
	stops int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	adj := make([][]pair, n)

	for _, flight := range flights {
		src, dst, cost := flight[0], flight[1], flight[2]
		adj[src] = append(adj[src], pair{dst: dst, cost: cost, stops: 0})
	}

	minHeap := &MinHeap{}
	heap.Push(minHeap, pair{dst: src, cost: 0, stops: 0})
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt64
	}

	for minHeap.Len() > 0 {
		pair := heap.Pop(minHeap).(pair)

		if pair.stops <= k {
			for _, next := range adj[pair.dst] {
				next.stops = pair.stops + 1
				next.cost += pair.cost
				if next.cost < distance[next.dst] {
					heap.Push(minHeap, next)
					distance[next.dst] = next.cost
				}
			}
		}
	}

	if distance[dst] == math.MaxInt64 {
		return -1
	}
	return distance[dst]
}

type MinHeap []pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i int, j int) bool {
	if h[i].stops == h[j].stops {
		return h[i].cost < h[j].cost
	}
	return h[i].stops < h[j].stops
}

func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(a interface{}) {
	*h = append(*h, a.(pair))
}

func (h *MinHeap) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
