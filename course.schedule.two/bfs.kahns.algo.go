package main

func findOrderBFS(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)

	adjacencyList := make(map[int][]int)

	for i := range prerequisites {
		adjacencyList[prerequisites[i][1]] = append(adjacencyList[prerequisites[i][1]], prerequisites[i][0])
		inDegree[prerequisites[i][0]]++
	}

	var q []int
	visited := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
			visited[i] = true
		}
	}

	var result []int
	for len(q) > 0 {
		course := q[0]
		q = q[1:]
		result = append(result, course)

		for _, nei := range adjacencyList[course] {
			if !visited[nei] {
				inDegree[nei]--
				if inDegree[nei] == 0 {
					if !visited[nei] {
						q = append(q, nei)
						visited[nei] = true
					}
				}
			}
		}
	}

	if len(result) == numCourses {
		return result
	}

	return []int{}
}
