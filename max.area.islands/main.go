package main

var Directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func inBoundary(i, j, m, n int) bool {
	if (0 <= i && i < m) && (0 <= j && j < n) {
		return true
	}
	return false
}

func dfs(i, j, m, n int, grid [][]int) int {
	if grid[i][j] == '0' {
		return 0
	}
	grid[i][j] = 0

	ans := 1
	for _, d := range Directions {
		row, col := i+d[0], j+d[1]
		if inBoundary(row, col, m, n) && grid[row][col] == 1 {
			ans += dfs(row, col, m, n, grid)
		}
	}
	return ans
}

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count := dfs(i, j, m, n, grid)
				if maxArea < count {
					maxArea = count
				}
			}
		}
	}

	return maxArea
}
