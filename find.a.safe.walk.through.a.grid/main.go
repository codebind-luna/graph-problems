package main

type pair struct {
	r, c, h int
}

func findSafeWalk(grid [][]int, health int) bool {
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	var q []pair

	p := pair{r: 0, c: 0, h: health}
	if grid[0][0] == 1 {
		p.h -= 1
	}
	q = append(q, p)
	visited[0][0] = true

	for len(q) > 0 {
		e := q[0]
		q = q[1:]

		r, c, h := e.r, e.c, e.h

		if r == rows-1 && c == cols-1 {
			return h > 0
		}

		if h > 0 {
			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if (0 <= nr && nr < rows && 0 <= nc && nc < cols) && !visited[nr][nc] {
					nh := h - grid[nr][nc]
					visited[nr][nc] = true
					q = append(q, pair{r: nr, c: nc, h: nh})
				}
			}
		}
	}
	return false
}
