package main

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	indegree := make([]int, n+1)
	adjList := make([][]int, n+1)

	for _, r := range relations {
		adjList[r[0]] = append(adjList[r[0]], r[1])
		indegree[r[1]]++
	}
	q := []int{}

	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	if len(q) == 0 {
		return -1
	}

	sem := 0

	for len(q) > 0 {
		n := len(q)
		n = min(n, k)
		sem++

		for i := 0; i < n; i++ {
			c := q[0]
			q = q[1:]

			for _, nei := range adjList[c] {
				indegree[nei]--
				if indegree[nei] == 0 {
					q = append(q, nei)
				}
			}
		}
	}

	return sem
}
