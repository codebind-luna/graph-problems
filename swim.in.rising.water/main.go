package main

import (
	"container/heap"
	"math"
)

type pair struct {
	row, col, t int
}

type MinHeap []*pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].t < h[j].t
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	timeMatrix := make([][]int, n)

	for i := 0; i < n; i++ {
		timeMatrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			timeMatrix[i][j] = math.MaxInt
		}
	}

	timeMatrix[0][0] = 0
	var h MinHeap
	heap.Init(&h)

	heap.Push(&h, &pair{row: 0, col: 0, t: grid[0][0]})

	for h.Len() > 0 {
		pos := heap.Pop(&h).(*pair)

		r, c, e := pos.row, pos.col, pos.t

		for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
			nr, nc := r+d[0], c+d[1]
			if 0 <= nr && nr < n && 0 <= nc && nc < n {
				ne := max(e, grid[nr][nc])
				if ne < timeMatrix[nr][nc] {
					timeMatrix[nr][nc] = ne
					heap.Push(&h, &pair{row: nr, col: nc, t: ne})
				}
			}

		}
	}

	return timeMatrix[n-1][n-1]
}
