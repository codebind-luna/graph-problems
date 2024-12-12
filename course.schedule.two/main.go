package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)

	adjacencyList := make(map[int][]int)

	for i := range prerequisites {
		adjacencyList[prerequisites[i][1]] = append(adjacencyList[prerequisites[i][1]], prerequisites[i][0])
		inDegree[prerequisites[i][0]]++
	}

	var q []int
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	var result []int
	for len(q) > 0 {
		course := q[0]
		q = q[1:]
		result = append(result, course)

		for _, nei := range adjacencyList[course] {
			inDegree[nei]--
			if inDegree[nei] == 0 {
				q = append(q, nei)
			}
		}

	}

	if len(result) == numCourses {
		return result
	}

	return []int{}
}
