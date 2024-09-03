package main

import "fmt"

func minimumSemesters(n int, relations [][]int) int {

	indegree := make([]int, n+1)

	fmt.Println(indegree)

	adjList := map[int][]int{}

	for _, rel := range relations {
		prev, next := rel[0], rel[1]

		adjList[prev] = append(adjList[prev], next)

		indegree[next] += 1
	}

	fmt.Println(indegree)
	fmt.Println()

	q := []int{}

	for idx, degree := range indegree {
		if idx == 0 {
			continue
		}

		if degree == 0 {
			q = append(q, idx)
		}
	}

	if len(q) == 0 {
		return -1
	}

	res, total := 0, 0

	for len(q) != 0 {

		n := len(q)

		total += n

		for i := 0; i < n; i++ {

			cur := q[0]
			q = q[1:]

			for _, nei := range adjList[cur] {
				indegree[nei] -= 1
				if indegree[nei] == 0 {
					q = append(q, nei)
				}
			}

		}

		res += 1
	}

	if total != n {
		return -1
	}

	return res
}
