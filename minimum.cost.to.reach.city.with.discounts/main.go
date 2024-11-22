package main

import (
	"container/heap"
	"math"
)

type Item struct {
	cost     int
	node     int
	discount int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func minimumCost(n int, highways [][]int, discounts int) int {
	graph := make([][][]int, n)
	for _, x := range highways {
		a, b, c := x[0], x[1], x[2]
		graph[a] = append(graph[a], []int{b, c})
		graph[b] = append(graph[b], []int{a, c})
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{cost: 0, node: 0, discount: 0})

	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, discounts+1)
		for j := range visited[i] {
			visited[i][j] = math.MaxInt32
		}
	}
	visited[0][0] = 0

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item)
		cost, city, dis := cur.cost, cur.node, cur.discount

		if city == n-1 {
			return cost
		}

		for _, x := range graph[city] {
			next, weight := x[0], x[1]
			if cost+weight < visited[next][dis] {
				heap.Push(pq, &Item{cost: cost + weight, node: next, discount: dis})
				visited[next][dis] = cost + weight
			}
			if dis < discounts && cost+weight/2 < visited[next][dis+1] {
				heap.Push(pq, &Item{cost: cost + weight/2, node: next, discount: dis + 1})
				visited[next][dis+1] = cost + weight/2
			}
		}
	}
	return -1
}
