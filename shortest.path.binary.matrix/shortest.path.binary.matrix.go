package main

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
	ROW := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	COL := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	queue := []Pair{{row: 0, col: 0, val: 1}}
	visited[0][0] = true
	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]

		if pair.row == n-1 && pair.col == n-1 {
			return pair.val
		}
		for j := 0; j < 8; j++ {
			if pair.row+ROW[j] >= 0 && pair.row+ROW[j] < n && pair.col+COL[j] >= 0 && pair.col+COL[j] < n &&
				grid[pair.row+ROW[j]][pair.col+COL[j]] == 0 && !visited[pair.row+ROW[j]][pair.col+COL[j]] {
				visited[pair.row+ROW[j]][pair.col+COL[j]] = true
				queue = append(queue, Pair{row: pair.row + ROW[j], col: pair.col + COL[j], val: pair.val + 1})
			}
		}

	}
	return -1
}

type Pair struct {
	row int
	col int
	val int
}
