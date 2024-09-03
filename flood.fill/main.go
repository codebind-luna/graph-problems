package main

var Directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func inBoundary(i, j, m, n int) bool {
	if (0 <= i && i < m) && (0 <= j && j < n) {
		return true
	}
	return false
}

func dfs(i, j, m, n int, image *[][]int, oC, nC int) {
	if (*image)[i][j] == nC {
		return
	}
	(*image)[i][j] = nC

	for _, d := range Directions {
		row, col := i+d[0], j+d[1]
		if inBoundary(row, col, m, n) && (*image)[row][col] == oC {
			dfs(row, col, m, n, image, oC, nC)
		}
	}
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	m := len(image)
	n := len(image[0])

	oC := image[sr][sc]
	dfs(sr, sc, m, n, &image, oC, color)
	return image
}
