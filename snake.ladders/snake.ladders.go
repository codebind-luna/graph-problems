package main

type entry struct {
	s int
	m int
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	queue := []entry{{1, 0}}

	posMap := make(map[int][2]int)

	start := 1
	flag := 0
	for i := n - 1; i >= 0; i-- {
		if flag%2 == 0 {
			for j := 0; j < n; j++ {
				posMap[start] = [2]int{i, j}
				start++
			}

		} else {
			for j := n - 1; j >= 0; j-- {
				posMap[start] = [2]int{i, j}
				start++
			}
		}
		flag++
	}

	visited := make(map[int]bool)

	visited[1] = true

	max := n * n
	// minimumMoves := 0
	for len(queue) > 0 {
		square := queue[0]
		queue = queue[1:]

		for i := 1; i < 7; i++ {
			nextS := square.s + i
			if square.s+i > max {
				nextS = max
			}

			row := posMap[nextS][0]
			col := posMap[nextS][1]

			if board[row][col] != -1 {
				nextS = board[row][col]
			}

			if nextS == max {
				return square.m + 1
			}
			if !visited[nextS] {
				queue = append(queue, entry{nextS, square.m + 1})
				visited[nextS] = true
			}
		}
	}

	return -1
}

func main() {
	board := [][]int{{-1, -1}, {-1, 3}, {-1, -1}}
	snakesAndLadders(board)
}
