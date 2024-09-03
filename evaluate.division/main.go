package main

type Vertex struct {
	Name  string
	Value float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	vertexMap := make(map[string][]Vertex)

	for idx, equation := range equations {
		v1 := equation[0]
		v2 := equation[1]

		vertexMap[v1] = append(vertexMap[v1], Vertex{v2, values[idx]})
		vertexMap[v2] = append(vertexMap[v2], Vertex{v1, 1.0 / values[idx]})
	}
	res := make([]float64, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		if _, ok := vertexMap[start]; !ok {
			res[i] = -1.0
			continue
		}
		visited := make(map[string]bool)
		var q []Vertex
		q = append(q, Vertex{Name: start, Value: 1.0})
		visited[start] = true

		ans := -1.0
		for len(q) > 0 {
			e := q[0]
			q = q[1:]
			n, w := e.Name, e.Value
			if n == end {
				ans = w
				break
			}
			for _, nei := range vertexMap[n] {
				if !visited[nei.Name] {
					q = append(q, Vertex{nei.Name, nei.Value * w})
					visited[nei.Name] = true
				}
			}
		}
		res[i] = ans
	}

	return res
}
