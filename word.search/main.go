package main

var Directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func inBoundary(i, j, m, n int) bool {
	if (0 <= i && i < m) && (0 <= j && j < n) {
		return true
	}
	return false
}

func dfs(i, j, m, n int, board [][]byte, idx int, visited map[pair]bool, word string) bool {
	if idx == len(word) {
		return true
	}

	if board[i][j] != word[idx] {
		return false
	}

	visited[pair{i, j}] = true
	ans := false
	for _, d := range Directions {
		row, col := i+d[0], j+d[1]
		if inBoundary(row, col, m, n) && !visited[pair{row, col}] {
			ans = ans || dfs(row, col, m, n, board, idx+1, visited, word)
			if ans {
				break
			}
		}
	}
	visited[pair{i, j}] = true
	return ans
}

type pair struct {
	row int
	col int
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	visited := make(map[pair]bool)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, m, n, board, 0, visited, word) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
	}
	exist(board, "SEE")
}
