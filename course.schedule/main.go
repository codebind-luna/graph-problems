package main

func dfs(n int, visited []bool, adj map[int][]int, s *[]int) {
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

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int)
	for _, pr := range prerequisites {
		adj[pr[1]] = append(adj[pr[1]], pr[0])
	}

	var s []int
	visited := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			dfs(i, visited, adj, &s)
		}
	}

	idx := 0
	m := make(map[int]int)

	for len(s) > 0 {
		node := s[len(s)-1]
		s = s[:len(s)-1]
		m[node] = idx
		idx++
	}

	for i := 0; i < numCourses; i++ {
		for _, nei := range adj[i] {
			if m[i] >= m[nei] {
				return false
			}
		}
	}
	return true
}
