package main

import "math"

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	costMatrix := make([][]int, 26)

	for i := 0; i < 26; i++ {
		costMatrix[i] = make([]int, 26)
	}

	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i != j {
				costMatrix[i][j] = math.MaxInt
			}
		}
	}

	for i := 0; i < len(original); i++ {
		costMatrix[original[i]-'a'][changed[i]-'a'] = min(costMatrix[original[i]-'a'][changed[i]-'a'], cost[i])
	}

	for via := 0; via < 26; via++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if costMatrix[i][via] == math.MaxInt || costMatrix[via][j] == math.MaxInt {
					continue
				}

				if i != j {
					costMatrix[i][j] = min(costMatrix[i][j], costMatrix[i][via]+costMatrix[via][j])
				}
			}
		}
	}

	totalCost := 0

	for i := 0; i < len(source); i++ {
		if costMatrix[source[i]-'a'][target[i]-'a'] == math.MaxInt {
			return -1
		}
		if source[i] == target[i] {
			continue
		}
		totalCost += costMatrix[source[i]-'a'][target[i]-'a']
	}

	return int64(totalCost)
}
