package main

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	ingredientsMap := map[string]bool{} // we treat ingredients and recipes the same (things we that exist or we can make basically)
	adjList := constructAdjList(recipes, ingredients)
	degrees := map[string]int{}

	for _, supply := range supplies {
		ingredientsMap[supply] = true
	}

	for source, dests := range adjList {
		_, exists := degrees[source]
		if !exists {
			degrees[source] = 0
		}
		for _, dest := range dests {
			// _, exists := degrees[dest]
			// if !exists {
			// 	degrees[dest] = 0
			// }

			degrees[dest]++
		}
	}

	q := []string{}

	for node, degree := range degrees {
		_, exists := ingredientsMap[node]
		if degree == 0 && exists {
			q = append(q, node)
		}
	}

	for len(q) > 0 {
		currNode := q[0]
		q = q[1:]

		for _, adjNode := range adjList[currNode] {
			degrees[adjNode]--
			if degrees[adjNode] == 0 {
				ingredientsMap[currNode] = true
				q = append(q, adjNode)
			}
		}
	}

	possibleRecipes := []string{}

	for _, recipe := range recipes {
		_, exists := ingredientsMap[recipe]
		if exists {
			possibleRecipes = append(possibleRecipes, recipe)
		}
	}

	return possibleRecipes
}

func constructAdjList(recipes []string, ingredients [][]string) map[string][]string {
	adjList := map[string][]string{}

	for i, source := range recipes {
		for _, dest := range ingredients[i] {
			// _, exists := adjList[dest]
			// if !exists {
			// 	adjList[dest] = []string{}
			// }

			adjList[dest] = append(adjList[dest], source)
		}
	}

	return adjList
}

func main() {
	recipes := []string{"bread"}
	ingredients := [][]string{{"yeast", "flour"}}

	supplies := []string{"yeast", "flour", "corn"}

	// log.Println(constructAdjList(recipes, ingredients))

	findAllRecipes(recipes, ingredients, supplies)
}
