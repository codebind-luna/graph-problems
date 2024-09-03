package main

import "sort"

func getAncestors(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	inDeg := make([]int, n)

	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		inDeg[e[1]]++
	}

	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if inDeg[i] == 0 {
			q = append(q, i)
		}
	}

	ancestors := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		ancestors[i] = make(map[int]bool)
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range graph[u] {
			inDeg[v]--
			ancestors[v][u] = true
			for ancestor := range ancestors[u] {
				ancestors[v][ancestor] = true
			}
			if inDeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, 0, len(ancestors[i]))
		for ancestor := range ancestors[i] {
			ans[i] = append(ans[i], ancestor)
		}
		sort.Ints(ans[i])
	}

	return ans
}
