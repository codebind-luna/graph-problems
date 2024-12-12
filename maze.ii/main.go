package main

import (
	"container/heap"
	"math"
)

type data struct {
	r, c, d int
}

type minHeap []data

func (mh minHeap) Len() int {
	return len(mh)
}

func (mh minHeap) Less(i, j int) bool {
	return mh[i].d < mh[j].d
}
func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.(data))
}
func (mh *minHeap) Pop() interface{} {
	x := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return x
}

func shortestDistance(maze [][]int, start []int, destination []int) int {
	rows, cols := len(maze), len(maze[0])
	distance := make([][]int, rows)

	for i := 0; i < rows; i++ {
		distance[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			distance[i][j] = math.MaxInt
		}
	}

	var pq minHeap
	heap.Init(&pq)
	heap.Push(&pq, data{r: start[0], c: start[1], d: 0})

	for pq.Len() > 0 {
		e := heap.Pop(&pq).(data)
		r, c, d := e.r, e.c, e.d

		for _, dir := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nr, nc := r+dir[0], c+dir[1]

			count := 0
			for nr >= 0 && nr < rows && nc >= 0 && nc < cols && maze[nr][nc] == 0 {
				count++
				nr, nc = nr+dir[0], nc+dir[1]
			}
			nd := d + count
			if nd < distance[nr-dir[0]][nc-dir[1]] {
				distance[nr-dir[0]][nc-dir[1]] = nd
				heap.Push(&pq, data{r: nr - dir[0], c: nc - dir[1], d: nd})
			}
		}
	}

	if distance[destination[0]][destination[1]] == math.MaxInt {
		return -1
	}

	return distance[destination[0]][destination[1]]
}
