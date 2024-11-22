package main

var Directions = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func inBoundry(i, j, m, n int) bool {
	if 0 <= i && i < m && 0 <= j && j < n {
		return true
	}
	return false
}

func dfs(rows, cols, row, col int, grid [][]int, seen [][]bool) {
	if !inBoundry(row, col, rows, cols) || grid[row][col] == 1 || seen[row][col] {
		return
	}
	seen[row][col] = true

	for _, d := range Directions {
		newRow, newCol := row+d[0], col+d[1]
		dfs(rows, cols, newRow, newCol, grid, seen)
	}
}

func closedIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	seen := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		seen[i] = make([]bool, cols)
	}

	// top edge
	for i := 0; i < cols; i++ {
		if grid[0][i] == 0 && !seen[0][i] {
			dfs(rows, cols, 0, i, grid, seen)
		}
	}

	// left edge
	for i := 0; i < rows; i++ {
		if grid[i][0] == 0 && !seen[i][0] {
			dfs(rows, cols, i, 0, grid, seen)
		}
	}

	// right edge
	for i := 0; i < rows; i++ {
		if grid[i][cols-1] == 0 && !seen[i][cols-1] {
			dfs(rows, cols, i, cols-1, grid, seen)
		}
	}

	// bottom edge
	for i := 0; i < cols; i++ {
		if grid[rows-1][i] == 0 && !seen[rows-1][i] {
			dfs(rows, cols, rows-1, i, grid, seen)
		}
	}

	noOfClosedIsland := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 && !seen[i][j] {
				dfs(rows, cols, i, j, grid, seen)
				noOfClosedIsland++
			}
		}
	}
	return noOfClosedIsland
}
