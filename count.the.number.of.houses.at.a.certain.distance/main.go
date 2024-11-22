package main

import "math"

func countOfPairs(n int, x int, y int) []int {
	distMatrix := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		distMatrix[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i != j {
				distMatrix[i][j] = math.MaxInt
			}
		}
	}

	for i := 1; i < n; i++ {
		distMatrix[i][i+1] = 1
		distMatrix[i+1][i] = 1
	}

	distMatrix[x][y], distMatrix[y][x] = 1, 1

	for via := 1; via <= n; via++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if distMatrix[i][via] == math.MaxInt || distMatrix[via][j] == math.MaxInt {
					continue
				}
				if i != j {
					distMatrix[i][j] = min(distMatrix[i][j], distMatrix[i][via]+distMatrix[via][j])
				}
			}
		}
	}

	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i != j {
				val := distMatrix[i][j]
				if 1 <= val && val <= n {
					result[val]++
				}
			}
		}
	}
	return result[1:]
}
