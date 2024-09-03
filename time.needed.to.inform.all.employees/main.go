package main

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	visited := make([]bool, n)

	adjList := make(map[int][]int)

	for i := range manager {
		adjList[manager[i]] = append(adjList[manager[i]], i)
	}

	var q []int
	q = append(q, headID)
	visited[headID] = true

	totalTime := 0
	for len(q) > 0 {
		s := len(q)
		maxTime := 0
		for i := 0; i < s; i++ {
			e := q[0]
			q = q[1:]
			maxTime = max(maxTime, informTime[e])
			for _, nei := range adjList[e] {
				if !visited[nei] {
					visited[nei] = true
					q = append(q, nei)
				}
			}
		}

		totalTime += maxTime
	}

	return totalTime
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
