/*
	Leetcode problem: https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/
*/

package main

import (
	"container/heap"
	"math"
)

type pair struct {
	r, c, d int
}

type minHeap []pair

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i int, j int) bool {
	return h[i].d < h[j].d
}

func (h minHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(a interface{}) {
	*h = append(*h, a.(pair))
}

func (h *minHeap) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func minimumObstacles(grid [][]int) int {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	rows, cols := len(grid), len(grid[0])
	distance := make([][]int, rows)
	for i := 0; i < rows; i++ {
		distance[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			distance[i][j] = math.MaxInt
		}
	}

	var h minHeap
	heap.Init(&h)
	heap.Push(&h, pair{r: 0, c: 0, d: 0})
	distance[0][0] = 0

	for h.Len() > 0 {
		e := heap.Pop(&h).(pair)
		r, c, dis := e.r, e.c, e.d

		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]

			if 0 <= nr && nr < rows && 0 <= nc && nc < cols {
				nd := dis + grid[nr][nc]

				if nd < distance[nr][nc] {
					distance[nr][nc] = nd
					heap.Push(&h, pair{r: nr, c: nc, d: nd})
				}
			}
		}
	}

	return distance[rows-1][cols-1]
}
