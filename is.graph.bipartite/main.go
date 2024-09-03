package main

func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if colors[i] != 0 {
			continue
		}

		queue := []int{i}
		colors[i] = 1
		for len(queue) > 0 {
			vertex := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[vertex] {
				if colors[neighbor] == 0 {
					colors[neighbor] = colors[vertex] * -1
					queue = append(queue, neighbor)
				} else if colors[neighbor] == colors[vertex] {
					return false
				}
			}
		}
	}

	return true
}
