package main

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	adjacencyList := make(map[int][]int)

	for i, v := range graph {
		adjacencyList[i] = v
	}

	queue := [][]int{{0}}

	for len(queue) > 0 {
		path := queue[0]

		lastNode := path[len(path)-1]
		queue = queue[1:]

		if lastNode == len(graph)-1 {
			res = append(res, path)
			continue
		}

		for _, adjacentNode := range adjacencyList[lastNode] {
			tmpPath := make([]int, len(path))
			// Copy needed in Golang because of slice implementation.
			// Without copy we will alse append values to the curPath.
			copy(tmpPath, path)
			tmpPath = append(tmpPath, adjacentNode)
			queue = append(queue, tmpPath)
		}

	}

	return res
}
