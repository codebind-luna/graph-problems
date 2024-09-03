package main

func validPath(edges [][]int, source int, destination int) bool {
	seen := make(map[int]bool)
	var q []int
	adjacencyList := make(map[int][]int)

	for _, v := range edges {
		adjacencyList[v[0]] = append(adjacencyList[v[0]], v[1])
		adjacencyList[v[1]] = append(adjacencyList[v[1]], v[0])
	}

	q = append(q, source)
	seen[source] = true

	for len(q) > 0 {
		top := q[0]
		if top == destination {
			return true
		}
		q = q[1:]
		for _, v := range adjacencyList[top] {
			if !seen[v] {
				q = append(q, v)
				seen[v] = true
			}
		}
	}
	return false
}
