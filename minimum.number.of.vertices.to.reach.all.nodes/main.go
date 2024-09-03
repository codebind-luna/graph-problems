package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	inDegrees := make([]int, n)

	for _, e := range edges {
		inDegrees[e[1]]++
	}

	var res []int
	for i := 0; i < n; i++ {
		if inDegrees[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}
