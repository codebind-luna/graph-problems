package main

func topoSort(conditions [][]int, k int) []int {
	inDegree := make([]int, k+1)
	adjList := make([][]int, k+1)

	for _, c := range conditions {
		adjList[c[0]] = append(adjList[c[0]], c[1])
		inDegree[c[1]]++
	}

	var q []int
	for i := 1; i <= k; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	ans := []int{}

	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		ans = append(ans, n)

		for _, nei := range adjList[n] {
			inDegree[nei]--
			if inDegree[nei] == 0 {
				q = append(q, nei)
			}
		}
	}

	return ans
}

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rows := topoSort(rowConditions, k)
	cols := topoSort(colConditions, k)

	if len(rows) != k || len(cols) != k {
		return [][]int{}
	}

	matrix := make([][]int, k)
	for i := 0; i < k; i++ {
		matrix[i] = make([]int, k)
	}

	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			if rows[i] == cols[j] {
				matrix[i][j] = rows[i]
			}
		}
	}

	return matrix
}
