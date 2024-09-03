package main

func dfs(node int, adjList map[int][]int, visited []bool) {
	if visited[node] {
		return
	}
	visited[node] = true

	for _, nei := range adjList[node] {
		if !visited[nei] {
			dfs(nei, adjList, visited)
		}
	}
}

func countComponents(n int, edges [][]int) int {
	adjList := make(map[int][]int)

	visited := make([]bool, n)
	for i := range edges {
		adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
		adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])
	}

	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			dfs(i, adjList, visited)
		}
	}

	return count
}
