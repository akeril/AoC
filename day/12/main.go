package main

import (
	"fmt"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	matrix := utils.ReadFile("input")
	visited := make([][]bool, 0)
	for _, row := range matrix {
		visited = append(visited, make([]bool, len(row)))
	}
	total := 0
	for i, row := range matrix {
		for j := range row {
			area, corners := dfs(matrix, visited, i, j, matrix[i][j])
			total += area * corners
		}
	}
	fmt.Println(total)
}

type move struct {
	x, y int
}

func dfs(matrix []string, visited [][]bool, i, j int, b byte) (int, int) {
	if !inBound(matrix, i, j) {
		return 0, 0
	}
	if matrix[i][j] != b {
		return 0, 0
	}
	if visited[i][j] {
		return 0, 0
	}
	visited[i][j] = true

	moves := []move{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	area := 1

	corners := countCorners(matrix, i, j, b)
	for _, move := range moves {
		a, c := dfs(matrix, visited, i+move.x, j+move.y, b)
		area += a
		corners += c
	}
	return area, corners
}

func inBound(matrix []string, i, j int) bool {
	return i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0])
}

func countCorners(matrix []string, i, j int, b byte) int {
	corners := 0
	c := matrix[i][j]
	// upper left
	if !same(matrix, i-1, j, c) && !same(matrix, i, j-1, c) {
		corners++
	}
	if same(matrix, i, j+1, c) && same(matrix, i+1, j, c) && !same(matrix, i+1, j+1, c) {
		corners++
	}
	// upper right
	if !same(matrix, i, j+1, c) && !same(matrix, i-1, j, c) {
		corners++
	}
	if same(matrix, i+1, j, c) && same(matrix, i, j-1, c) && !same(matrix, i+1, j-1, c) {
		corners++
	}
	// lower left
	if !same(matrix, i+1, j, c) && !same(matrix, i, j-1, c) {
		corners++
	}
	if same(matrix, i-1, j, c) && same(matrix, i, j+1, c) && !same(matrix, i-1, j+1, c) {
		corners++
	}
	// lower right
	if !same(matrix, i+1, j, c) && !same(matrix, i, j+1, c) {
		corners++
	}
	if same(matrix, i, j-1, c) && same(matrix, i-1, j, c) && !same(matrix, i-1, j-1, c) {
		corners++
	}
	return corners
}

func same(matrix []string, i, j int, b byte) bool {
	return inBound(matrix, i, j) && matrix[i][j] == b
}
