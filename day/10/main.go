package main

import (
	"fmt"

	"github.com/kjabin/aoc2024/utils"
)

var p = fmt.Sprintf

func main() {
	input := utils.ReadFile("input")
	matrix := utils.ToArr2D(input)
	count := 0
	for i, row := range matrix {
		for j := range row {
			visited := make(map[move]bool)
			dfs(matrix, i, j, '0', visited)
			count += len(visited)
		}
	}
	fmt.Println(count)
}

type move struct {
	x, y int
}

func dfs(matrix [][]rune, i, j int, c rune, mp map[move]bool) {
	if !inBound(matrix, i, j) {
		return
	}
	if matrix[i][j] != c {
		return
	}
	if matrix[i][j] == '9' {
		mp[move{i, j}] = true
		return
	}

	moves := []move{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, m := range moves {
		dfs(matrix, i+m.x, j+m.y, matrix[i][j]+1, mp)
	}
	return
}

func inBound(matrix [][]rune, i, j int) bool {
	return i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0])
}
