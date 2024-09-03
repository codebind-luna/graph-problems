package main

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	indegree := make([]int, n)
	adj := make([][]int, n)

	for i := range graph {
		for _, node := range graph[i] {
			adj[node] = append(adj[node], i)
			indegree[i]++
		}
	}

	var q []int
	// Push all the nodes with indegree zero in the queue.
	for i := range indegree {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	safe := make([]bool, n)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		safe[node] = true

		for _, neighbor := range adj[node] {
			// Delete the edge "node -> neighbor".
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	var safeNodes []int
	for i := range safe {
		if safe[i] {
			safeNodes = append(safeNodes, i)
		}
	}

	return safeNodes
}
