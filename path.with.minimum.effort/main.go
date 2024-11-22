package main

import (
	"container/heap"
	"math"
)

type pair struct {
	row, col, effort int
}

type MinHeap []*pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].effort < h[j].effort
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

func minimumEffortPath(heights [][]int) int {
	n := len(heights)
	m := len(heights[0])
	effortMatrix := make([][]int, n)

	for i := 0; i < n; i++ {
		effortMatrix[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			effortMatrix[i][j] = math.MaxInt
		}
	}

	effortMatrix[0][0] = 0
	var h MinHeap
	heap.Init(&h)

	heap.Push(&h, &pair{row: 0, col: 0, effort: 0})

	for h.Len() > 0 {
		pos := heap.Pop(&h).(*pair)

		r, c, e := pos.row, pos.col, pos.effort

		for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
			nr, nc := r+d[0], c+d[1]
			if 0 <= nr && nr < n && 0 <= nc && nc < m {
				ne := max(e, abs(heights[nr][nc]-heights[r][c]))
				if ne < effortMatrix[nr][nc] {
					effortMatrix[nr][nc] = ne
					heap.Push(&h, &pair{row: nr, col: nc, effort: ne})
				}
			}

		}
	}

	return effortMatrix[n-1][m-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
