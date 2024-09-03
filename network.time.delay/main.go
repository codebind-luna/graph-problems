package main

// type Pair struct {
// 	Key, Value int
// }

// func networkDelayTime(times [][]int, n, k int) int {
// 	// Adjacency list
// 	adj := make(map[int][]Pair)

// 	for _, time := range times {
// 		source, dest, travelTime := time[0], time[1], time[2]
// 		adj[source] = append(adj[source], Pair{travelTime, dest})
// 	}

// 	visited := make([]bool, n+1)

// 	maxT := BFS(visited, k, adj)

// 	for i := 1; i <= n; i++ {
// 		if !visited[i] {
// 			return -1
// 		}
// 	}

// 	return maxT
// }

// func BFS(visited []bool, sourceNode int, adj map[int][]Pair) int {
// 	var queue [][2]int
// 	queue = append(queue, [2]int{sourceNode, 0})
// 	visited[sourceNode] = true

// 	maxTime := 0
// 	for len(queue) > 0 {
// 		currNode := queue[0]
// 		queue = queue[1:]
// 		maxTime = max(maxTime, currNode[1])
// 		// Broadcast the signal to adjacent nodes
// 		for _, edge := range adj[currNode[0]] {
// 			time, neighborNode := edge.Key, edge.Value
// 			arrivalTime := currNode[1] + time
// 			if !visited[neighborNode] {
// 				visited[neighborNode] = true
// 				queue = append(queue, [2]int{neighborNode, arrivalTime})
// 			}
// 		}
// 	}
// 	return maxTime
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
