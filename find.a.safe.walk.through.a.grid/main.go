/*
	Leetcode problem: https://leetcode.com/problems/find-a-safe-walk-through-a-grid/
*/
package main

import (
	"container/heap"
	"math"
)

type pair struct {
	r, c, h int
}

type minHeap []pair

func (mh minHeap) Len() int {
	return len(mh)
}

func (mh minHeap) Less(i int, j int) bool {
	return mh[i].h > mh[j].h
}

func (mh minHeap) Swap(i int, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap) Push(a interface{}) {
	*mh = append(*mh, a.(pair))
}

func (mh *minHeap) Pop() interface{} {
	l := len(*mh)
	res := (*mh)[l-1]
	*mh = (*mh)[:l-1]
	return res
}

func findSafeWalk(grid [][]int, health int) bool {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	rows, cols := len(grid), len(grid[0])
	healthCnt := make([][]int, rows)
	for i := 0; i < rows; i++ {
		healthCnt[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			healthCnt[i][j] = math.MinInt
		}
	}

	var mh minHeap
	heap.Init(&mh)
	heap.Push(&mh, pair{r: 0, c: 0, h: health - grid[0][0]})
	healthCnt[0][0] = 0

	for mh.Len() > 0 {
		e := heap.Pop(&mh).(pair)
		r, c, h := e.r, e.c, e.h

		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]

			if 0 <= nr && nr < rows && 0 <= nc && nc < cols {
				nh := h - grid[nr][nc]

				if nh > healthCnt[nr][nc] {
					healthCnt[nr][nc] = nh
					heap.Push(&mh, pair{r: nr, c: nc, h: nh})
				}
			}
		}
	}

	return healthCnt[rows-1][cols-1] >= 1
}
