package main

type pair struct {
	r     int
	c     int
	dis   int
	eLeft int
}

type state struct {
	r     int
	c     int
	eLeft int
}

func shortestPath(grid [][]int, k int) int {
	rows, cols := len(grid), len(grid[0])

	if k >= rows+cols-2 {
		return rows + cols - 2
	}

	var q []pair
	q = append(q, pair{r: 0, c: 0, dis: 0, eLeft: k})
	seen := make(map[state]bool)
	seen[state{r: 0, c: 0, eLeft: k}] = true

	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		row, col, dis, eLeft := item.r, item.c, item.dis, item.eLeft

		if row == rows-1 && col == cols-1 {
			return dis
		}

		for _, direction := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			newRow, newCol := row+direction[0], col+direction[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
				newEliminations := eLeft - grid[newRow][newCol]

				if newEliminations >= 0 {
					if !seen[state{r: newRow, c: newCol, eLeft: newEliminations}] {
						seen[state{r: newRow, c: newCol, eLeft: newEliminations}] = true
						q = append(q, pair{r: newRow, c: newCol, dis: dis + 1, eLeft: newEliminations})
					}
				}
			}
		}
	}

	return -1
}
