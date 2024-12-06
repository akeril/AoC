package main

import (
	"fmt"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	matrix := utils.ToArr2D(input)
	fmt.Println(calcLoops(matrix))
}

func calcLoops(matrix [][]rune) int {
	count := 0
	x, y := startPos(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != '.' {
				continue
			}
			matrix[i][j] = '#'
			visited := make(map[Position]bool)
			if !walk(matrix, visited, x, y, matrix[x][y]) {
				count++
			}
			matrix[i][j] = '.'
		}
	}
	return count
}

func calcSteps(matrix [][]rune) int {
	x, y := startPos(matrix)
	visited := make(map[Position]bool)
	walk(matrix, visited, x, y, matrix[x][y])
	return len(visited)
}

func walk(matrix [][]rune, visited map[Position]bool, x, y int, direction rune) bool {
	pos := Position{x: x, y: y, direction: direction}
	if _, ok := visited[pos]; ok {
		return false
	}
	visited[pos] = true
	dx, dy := calcMove(direction)
	if !inBound(matrix, x+dx, y+dy) {
		return true
	}
	if matrix[x+dx][y+dy] != '#' {
		return walk(matrix, visited, x+dx, y+dy, direction)
	}
	for matrix[x+dx][y+dy] == '#' {
		direction = turn(direction)
		dx, dy = calcMove(direction)
	}
	return walk(matrix, visited, x+dx, y+dy, direction)
}

func startPos(matrix [][]rune) (int, int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != '.' && matrix[i][j] != '#' {
				return i, j
			}
		}
	}
	// unreachable
	return -1, -1
}

func turn(direction rune) rune {
	if direction == '>' {
		return 'V'
	}
	if direction == '<' {
		return '^'
	}
	if direction == '^' {
		return '>'
	}
	if direction == 'V' {
		return '<'
	}
	// unreachable
	return '0'
}
func calcMove(direction rune) (int, int) {
	if direction == '>' {
		return 0, 1
	}
	if direction == '<' {
		return 0, -1
	}
	if direction == '^' {
		return -1, 0
	}
	if direction == 'V' {
		return 1, 0
	}
	// unreachable
	return -1, -1
}

func inBound(matrix [][]rune, i, j int) bool {
	return i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0])
}

type Position struct {
	direction rune
	x, y      int
}
