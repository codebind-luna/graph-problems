package main

func dfs(n int, visited []bool, adj [][]int, s *[]int) {
	if visited[n] {
		return
	}
	visited[n] = true

	for _, nei := range adj[n] {
		if !visited[nei] {
			dfs(nei, visited, adj, s)
		}
	}
	*s = append(*s, n)
}

func isCycle(V int, adj [][]int) bool {
	var s []int
	visited := make([]bool, V)

	for i := 0; i < V; i++ {
		if !visited[i] {
			dfs(i, visited, adj, &s)
		}
	}

	idx := 0
	m := make(map[int]int)

	for _, v := range s {
		m[v] = idx
		idx++
	}

	for i := 0; i < V; i++ {
		for _, nei := range adj[i] {
			if m[i] >= m[nei] {
				return true
			}
		}
	}
	return false
}
