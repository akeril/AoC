package main

import (
	"fmt"

	"github.com/akeril/aoc2024/utils"
)

var p = fmt.Sprintf

func main() {
	input := utils.ReadFile("input")
	matrix := utils.ToArr2D(input)
	count := 0
	for i, row := range matrix {
		for j := range row {
			trails := dfs(matrix, i, j, '0')
			count += trails
		}
	}
	fmt.Println(count)
}

type move struct {
	x, y int
}

func dfs(matrix [][]rune, i, j int, c rune) int {
	if !inBound(matrix, i, j) {
		return 0
	}
	if matrix[i][j] != c {
		return 0
	}
	if matrix[i][j] == '9' {
		return 1
	}

	total := 0
	moves := []move{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, m := range moves {
		total += dfs(matrix, i+m.x, j+m.y, matrix[i][j]+1)
	}
	return total
}

func inBound(matrix [][]rune, i, j int) bool {
	return i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0])
}
