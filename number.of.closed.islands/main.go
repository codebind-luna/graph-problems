package main

var Directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func inBoundary(i, j, m, n int) bool {
	if (0 <= i && i < m) && (0 <= j && j < n) {
		return true
	}
	return false
}

func dfs(i, j, m, n int, grid *[][]int) {
	if (*grid)[i][j] == 1 {
		return
	}
	(*grid)[i][j] = 1

	for _, d := range Directions {
		row, col := i+d[0], j+d[1]
		if inBoundary(row, col, m, n) && (*grid)[row][col] == 0 {
			dfs(row, col, m, n, grid)
		}
	}
}

func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// flood fill left corner and right corner
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			dfs(i, 0, m, n, &grid)
		}
		if grid[i][n-1] == 0 {
			dfs(i, n-1, m, n, &grid)
		}
	}

	// flood fill top corner and bottom corner
	for i := 0; i < n; i++ {
		if grid[0][i] == 0 {
			dfs(0, i, m, n, &grid)
		}
		if grid[m-1][i] == 0 {
			dfs(m-1, i, m, n, &grid)
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dfs(i, j, m, n, &grid)
				count++
			}
		}
	}
	return count
}
