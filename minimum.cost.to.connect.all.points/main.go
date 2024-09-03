package main

import (
	"container/heap"
	"math"
)

type entry struct {
	wt, node int
}

type priorityQueue []entry

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].wt < pq[j].wt
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(entry))
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func minCostConnectPoints(points [][]int) int {
	sum := 0
	n := len(points)
	visited := make([]bool, n)

	heapDict := make([]int, n)
	for i := 0; i < n; i++ {
		heapDict[i] = math.MaxInt // Initialize all distances to infinity
	}
	heapDict[0] = 0 // Start node

	var h priorityQueue
	heap.Init(&h)
	heap.Push(&h, entry{0, 0})

	for len(h) > 0 {
		e := heap.Pop(&h).(entry)
		w, no := e.wt, e.node

		if visited[no] || heapDict[no] < w {
			continue
		}
		visited[no] = true
		sum += w

		for i := 0; i < n; i++ {
			if !visited[i] {
				nw := abs(points[i][0]-points[no][0]) + abs(points[i][1]-points[no][1])
				if nw < heapDict[i] {
					heap.Push(&h, entry{nw, i})
					heapDict[i] = nw
				}
			}
		}
	}

	return sum
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
