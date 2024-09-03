package main

const (
	Fresh  = 1
	Rotten = 2
)

type Pair struct {
	row int
	col int
	val int
}

// var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ROW := []int{-1, 0, 0, 1}
	COL := []int{0, -1, 1, 0}

	var rotten []Pair
	fresh := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case Fresh:
				fresh++
			case Rotten:
				rotten = append(rotten, Pair{i, j, 0})
			}
		}
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for _, v := range rotten {
		visited[v.row][v.col] = true
	}

	time := 0
	for len(rotten) > 0 {
		size := len(rotten)
		time++
		for i := 0; i < size; i++ {
			pair := rotten[0]
			rotten = rotten[1:]

			q := []Pair{}
			for j := 0; j < 4; j++ {
				if pair.row+ROW[j] >= 0 && pair.row+ROW[j] < m && pair.col+COL[j] >= 0 && pair.col+COL[j] < n &&
					grid[pair.row+ROW[j]][pair.col+COL[j]] == Fresh && !visited[pair.row+ROW[j]][pair.col+COL[j]] {
					visited[pair.row+ROW[j]][pair.col+COL[j]] = true
					grid[pair.row+ROW[j]][pair.col+COL[j]] = Rotten
					fresh--
					q = append(q, Pair{row: pair.row + ROW[j], col: pair.col + COL[j], val: pair.val + 1})
				}
			}
			rotten = append(rotten, q...)
		}
	}

	if fresh == 0 {
		return time
	}

	return -1
}
