package main

type pair struct {
	r int
	c int
	d int
	e int
}

type state struct {
	r int
	c int
	e int
}

func shortestPath(grid [][]int, k int) int {
	rows, cols := len(grid), len(grid[0])

	var q []pair
	q = append(q, pair{r: 0, c: 0, d: 0, e: 0})
	seen := make(map[state]bool)
	seen[state{r: 0, c: 0, e: 0}] = true

	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		row, col, dis, e := item.r, item.c, item.d, item.e

		if row == rows-1 && col == cols-1 {
			return dis
		}

		for _, direction := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			newRow, newCol := row+direction[0], col+direction[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
				newEliminations := e + grid[newRow][newCol]

				if newEliminations <= k {
					if !seen[state{r: newRow, c: newCol, e: newEliminations}] {
						seen[state{r: newRow, c: newCol, e: newEliminations}] = true
						q = append(q, pair{r: newRow, c: newCol, d: dis + 1, e: newEliminations})
					}
				}
			}
		}
	}

	return -1
}
