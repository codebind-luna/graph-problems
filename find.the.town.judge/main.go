package main

func findJudge(n int, trust [][]int) int {
	inDegrees, outDegrees := make(map[int]int), make(map[int]int)

	for i := range trust {
		outDegrees[trust[i][0]]++
		inDegrees[trust[i][1]]++
	}

	for i := 1; i <= n; i++ {
		if inDegrees[i] == n-1 && outDegrees[i] == 0 {
			return i
		}
	}

	return -1
}
