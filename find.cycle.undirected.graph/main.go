package main

func isCycle(V int, adj [][]int) bool {
	visited := make([]bool, V)
	var q [][2]int

	for i := 0; i < V; i++ {
		if !visited[i] {
			q = append(q, [2]int{0, -1})
			visited[0] = true
		}

		for len(q) > 0 {
			e := q[0]
			n, p := e[0], e[1]
			q = q[1:]

			for _, nei := range adj[n] {
				if p != nei && visited[nei] {
					return true
				}

				if p != nei && !visited[nei] {
					q = append(q, [2]int{nei, n})
					visited[nei] = true
				}
			}
		}
	}

	return false
}
