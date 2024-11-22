package main

var Directions = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func inBoundry(i, j, m, n int) bool {
	if 0 <= i && i < m && 0 <= j && j < n {
		return true
	}
	return false
}

func dfs(rows, cols, row, col int, grid [][]byte, seen [][]bool) {
	if !inBoundry(row, col, rows, cols) || grid[row][col] == '0' || seen[row][col] {
		return
	}
	seen[row][col] = true

	for _, d := range Directions {
		newRow, newCol := row+d[0], col+d[1]
		dfs(rows, cols, newRow, newCol, grid, seen)
	}
}

func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	seen := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		seen[i] = make([]bool, cols)
	}

	noOfIsland := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !seen[i][j] {
				dfs(rows, cols, i, j, grid, seen)
				noOfIsland++
			}
		}
	}
	return noOfIsland
}
