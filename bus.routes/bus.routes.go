package main

func numBusesToDestination(routes [][]int, source int, target int) int {
	graph := make(map[int][]int)

	for i, r := range routes {
		for _, stop := range r {
			graph[stop] = append(graph[stop], i)
		}
	}

	queue := []int{source}
	seenStop := make(map[int]bool)
	seenStop[source] = true

	seenRoute := make(map[int]bool)

	ans := 0
	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			stop := queue[0]
			queue = queue[1:]

			if stop == target {
				return ans
			}

			list := []int{}
			for _, route := range graph[stop] {
				if !seenRoute[route] {
					for _, option := range routes[route] {
						if !seenStop[option] {
							list = append(list, option)
							seenStop[option] = true
						}
					}
				}
				seenRoute[route] = true
			}

			queue = append(queue, list...)

		}
		ans++
	}
	return -1
}
