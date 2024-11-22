package main

type pair struct {
	r int
	c int
	d int
}

func getFood(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	var q []pair
	visited := make([][]bool, m)

	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '*' {
				q = append(q, pair{r: i, c: j, d: 0})
				visited[i][j] = true
				break
			}
		}
	}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]

		r, c, dis := e.r, e.c, e.d

		for grid[r][c] == '#' {
			return dis
		}

		for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nr, nc, nd := r+d[0], c+d[1], (dis + 1)

			if 0 <= nr && nr < m && 0 <= nc && nc < n && !visited[nr][nc] && grid[nr][nc] != 'X' {
				visited[nr][nc] = true
				q = append(q, pair{r: nr, c: nc, d: nd})
			}
		}
	}

	return -1
}
