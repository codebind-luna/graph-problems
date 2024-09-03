package main

func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		inDegrees := make([]int, n)
		outDegrees := make([]int, n)

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j && knows(i, j) {
					outDegrees[i]++
					inDegrees[j]++
				}
			}
		}

		for i := 0; i < n; i++ {
			if inDegrees[i] == n-1 && outDegrees[i] == 0 {
				return i
			}
		}
		return -1
	}
}
