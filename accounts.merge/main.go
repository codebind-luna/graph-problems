package main

import (
	"sort"
)

func dfs(node string, adjList map[string][]string, visited map[string]bool, stack *[]string) {
	if visited[node] {
		return
	}

	visited[node] = true
	*stack = append(*stack, node)

	for _, nei := range adjList[node] {
		if !visited[nei] {
			dfs(nei, adjList, visited, stack)
		}
	}
}

func accountsMerge(accounts [][]string) [][]string {
	adjList := make(map[string][]string)

	emailPersonMap := make(map[string]string)

	for _, account := range accounts {
		emails := account[1:]
		for j := 0; j < len(emails); j++ {
			emailPersonMap[emails[j]] = account[0]
			for k := 0; k < len(emails); k++ {
				if k != j {
					adjList[emails[j]] = append(adjList[emails[j]], emails[k])
				}
			}
		}
	}

	visited := make(map[string]bool)

	result := [][]string{}

	for k, v := range emailPersonMap {
		if !visited[k] {
			entry := []string{}
			dfs(k, adjList, visited, &entry)
			sort.Strings(entry)
			result = append(result, append([]string{v}, entry...))
		}
	}

	return result
}
