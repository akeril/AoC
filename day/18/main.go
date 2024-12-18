package main

import (
	"fmt"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	matrix := MakeArr(71, 71)
	for _, line := range input[:1024] {
		values := utils.ToIntArr(line, ",")
		x, y := values[0], values[1]
		matrix[x][y] = 1
	}

	for _, line := range input[1024:] {
		values := utils.ToIntArr(line, ",")
		x, y := values[0], values[1]
		matrix[x][y] = 1

		b := bfs(matrix)
		if !b {
			fmt.Println(line)
			break
		}
	}

}

func bfs(matrix [][]int) bool {
	visited := make(map[Position]bool)
	start := Position{x: 0, y: 0, d: 0}
	queue := []Position{start}

	moves := []Move{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if !inBound(matrix, p.x, p.y) {
			continue
		}
		if matrix[p.x][p.y] == 1 {
			continue
		}
		if visited[p.Value()] {
			continue
		}
		if p.x == 70 && p.y == 70 {
			return true
		}
		visited[p.Value()] = true
		for _, m := range moves {
			queue = append(queue, Position{p.x + m.dx, p.y + m.dy, p.d + 1})
		}
	}
	return false
}

type Position struct {
	x, y, d int
}

func (p Position) Value() Position {
	p.d = 0
	return p
}

type Move struct {
	dx, dy int
}

func MakeArr(n, m int) [][]int {
	var arr [][]int
	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < m; j++ {
			row = append(row, 0)
		}
		arr = append(arr, row)
	}
	return arr
}

func inBound(matrix [][]int, i, j int) bool {
	return i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0])
}
