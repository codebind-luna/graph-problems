package main

func dfs(n int, visited []bool, adjList map[int][][]int) int {
	if visited[n] {
		return 0
	}

	visited[n] = true
	cnt := 0

	for _, v := range adjList[n] {
		if !visited[v[0]] {
			cnt += v[1] + dfs(v[0], visited, adjList)
		}
	}

	return cnt
}

func minReorder(n int, connections [][]int) int {
	adjList := make(map[int][][]int)

	for _, e := range connections {
		adjList[e[0]] = append(adjList[e[0]], []int{e[1], 1})
		adjList[e[1]] = append(adjList[e[1]], []int{e[0], 0})
	}

	visited := make([]bool, n)
	// var q []int
	// q = append(q, 0)
	// visited[0] = true
	// ans := 0

	// for len(q) > 0 {
	// 	e := q[0]
	// 	q = q[1:]

	// 	for _, v := range adjList[e] {
	// 		if !visited[v[0]] {
	// 			visited[v[0]] = true
	// 			ans += v[1]
	// 			q = append(q, v[0])
	// 		}
	// 	}
	// }

	// return ans
	return dfs(0, visited, adjList)
}
