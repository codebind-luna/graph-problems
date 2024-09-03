package main

func dfs(node int, visited map[int]bool, adjacencyList map[int][]int, stack *[]int) {
	if visited[node] {
		return
	}

	visited[node] = true

	for _, v := range adjacencyList[node] {
		if !visited[v] {
			dfs(v, visited, adjacencyList, stack)
		}
	}

	*stack = append(*stack, node)
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjacencyList := map[int][]int{}
	stack := []int{}
	visited := map[int]bool{}

	for _, v := range prerequisites {
		adjacencyList[v[1]] = append(adjacencyList[v[1]], v[0])
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			dfs(i, visited, adjacencyList, &stack)
		}
	}

	var order []int

	position := map[int]int{}
	indx := 0

	for len(stack) > 0 {
		course := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, course)
		position[course] = indx
		indx++
	}

	for i := 0; i < numCourses; i++ {
		for _, v := range adjacencyList[i] {
			// If parent vertex
			// does not appear first
			// 5 => [5] for self dependency also cycle exists.
			if position[i] >= position[v] {
				return []int{}
			}
		}
	}

	return order
}
