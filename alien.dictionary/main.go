package main

import "strings"

func dfs(node string, visited map[string]bool, adjacencyList map[string]map[string]bool, stack *[]string) {
	if visited[node] {
		return
	}

	visited[node] = true

	for k, _ := range adjacencyList[node] {
		if !visited[k] {
			dfs(k, visited, adjacencyList, stack)
		}
	}

	*stack = append(*stack, node)
}

func alienOrder(words []string) string {
	adjacencyList := make(map[string]map[string]bool)

	for i := 0; i < len(words); i++ {
		for _, v := range words[i] {
			adjacencyList[string(v)] = make(map[string]bool)
		}
	}

	stack := []string{}
	visited := map[string]bool{}

	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]
		minLength := min(len(word1), len(word2))

		if len(word1) > len(word2) && word1[:minLength] == word2[:minLength] {
			return ""
		}

		for j := 0; j < minLength; j++ {
			if word1[j] != word2[j] {
				adjacencyList[string(word1[j])][string(word2[j])] = true
				break
			}
		}
	}

	for k, _ := range adjacencyList {
		if !visited[k] {
			dfs(k, visited, adjacencyList, &stack)
		}
	}

	var order []string

	position := map[string]int{}
	indx := 0

	for len(stack) > 0 {
		alpha := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, alpha)
		position[alpha] = indx
		indx++
	}

	for k := range adjacencyList {
		for i := range adjacencyList[k] {
			// If parent vertex
			// does not appear first
			// 5 => [5] for self dependency also cycle exists.
			if position[k] >= position[i] {
				return ""
			}
		}
	}
	return strings.Join(order, "")
}
