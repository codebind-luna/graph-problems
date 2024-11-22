package main

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	distMatrix := make([][]int, n)

	for i := 0; i < n; i++ {
		distMatrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				distMatrix[i][j] = 0
			} else {
				distMatrix[i][j] = math.MaxInt
			}
		}
	}

	for _, edge := range edges {
		distMatrix[edge[0]][edge[1]] = edge[2]
		distMatrix[edge[1]][edge[0]] = edge[2]
	}

	for via := 0; via < n; via++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if distMatrix[i][via] == math.MaxInt || distMatrix[via][j] == math.MaxInt {
					continue
				}

				if i != j {
					distMatrix[i][j] = min(distMatrix[i][j], distMatrix[i][via]+distMatrix[via][j])
				}
			}
		}
	}

	cityWithfewerNeighbors := -1
	minReachables := math.MaxInt

	for i := 0; i < n; i++ {
		reachablesWithinThreshold := 0
		for j := 0; j < n; j++ {
			if i != j && distMatrix[i][j] <= distanceThreshold {
				reachablesWithinThreshold++
			}
		}
		if reachablesWithinThreshold <= minReachables {
			minReachables = reachablesWithinThreshold
			cityWithfewerNeighbors = i
		}
	}
	return cityWithfewerNeighbors
}
