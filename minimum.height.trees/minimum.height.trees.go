package main

import "log"

func findMinHeightTrees(n int, edges [][]int) []int {
	// find a node with minimum distance from other nodes -> find the center node of the tree alike graph
	// can go from outside to inside, layer by layer. at last there could either 1 or 2 nodes left
	// i.e 1 - 2 - 3 , or 1 - 2 - 3 - 4
	// similar like topological sorting, use queue to help store current leaf nodes
	// time O(N), space O(N)
	// build graph with adjesent list, map[int]map[int]bool{}
	res := []int{}
	if len(edges) < 2 {
		for i := 0; i < n; i++ {
			res = append(res, i)
		}
		return res
	}
	graph := buildGraph(edges)
	queue := []int{}
	// find first layer leaf
	for l, neis := range graph {
		if len(neis) == 1 {
			// Note: save leaf
			queue = append(queue, l)
		}
	}
	remainingNode := n
	for remainingNode > 2 {
		curLeaf := len(queue)
		remainingNode -= curLeaf
		newLeaf := []int{}
		for _, leaf := range queue {
			lNeis := graph[leaf]
			for ln := range lNeis {
				lnNeis := graph[ln]
				delete(lnNeis, leaf)
				if len(lnNeis) == 1 {
					// NOTE: save leaf
					newLeaf = append(newLeaf, ln)
				}
			}
			delete(graph, leaf)
		}
		queue = newLeaf
	}
	return queue
}

func buildGraph(edges [][]int) map[int]map[int]bool {
	graph := make(map[int]map[int]bool)

	for i := 0; i < len(edges)+1; i++ {
		graph[i] = make(map[int]bool)
	}

	for _, e := range edges {
		n1 := e[0]
		n2 := e[1]

		graph[n1][n2] = true
		graph[n2][n1] = true
		// // set n1 : n2
		// neis, ok := graph[n1]
		// if !ok {
		// 	neis = map[int]bool{}
		// }
		// neis[n2] = true
		// graph[n1] = neis
		// // set n2 : n1
		// neis2, ok2 := graph[n2]
		// if !ok2 {
		// 	neis2 = map[int]bool{}
		// }
		// neis2[n1] = true
		// graph[n2] = neis2
	}
	return graph
}

func main() {
	edges := [][]int{{1, 0}, {1, 2}, {1, 3}}
	log.Println(buildGraph(edges))
}
