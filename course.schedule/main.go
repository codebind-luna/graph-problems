
package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacencyList := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, p := range prerequisites {
		adjacencyList[p[1]] = append(adjacencyList[p[1]], p[0])
		inDegree[p[0]]++
	}

	q := []int{}

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	finished := 0

	for len(q) > 0 {
		c := q[len(q)-1]
		q = q[:len(q)-1]

		finished++

		for _, n := range adjacencyList[c] {
			inDegree[n]--
			if inDegree[n] == 0 {
				q = append(q, n)
			}
		}
	}
	return finished == numCourses
}
