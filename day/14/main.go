package main

import (
	"fmt"
	"strings"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")

	var bots []Bot
	for _, line := range input {
		b := strings.Split(line, ",")
		bots = append(bots, Bot{
			utils.ToInt(b[1]),
			utils.ToInt(b[0]),
			utils.ToInt(b[3]),
			utils.ToInt(b[2]),
		})
	}
	t := 0
	h, w := 103, 101
	for {
		matrix := Matrix(h, w)
		for _, bot := range bots {
			px := mod(bot.x+t*bot.vx, h)
			py := mod(bot.y+t*bot.vy, w)
			matrix[px][py]++
		}
		fmt.Println(t)
		if largestComponent(matrix) > 100 {
			Print(matrix)
			var z string
			fmt.Scanln(&z)
		}
		t++
	}
}

func largestComponent(matrix [][]int) int {
	largest := 0
	visited := make(map[move]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			largest = max(largest, dfs(matrix, i, j, visited))
		}
	}
	return largest
}

type move struct {
	x, y int
}

func dfs(matrix [][]int, i, j int, visited map[move]bool) int {
	if !inBound(matrix, i, j) {
		return 0
	}
	if matrix[i][j] == 0 {
		return 0
	}
	if visited[move{i, j}] {
		return 0
	}
	visited[move{i, j}] = true
	count := 1
	moves := []move{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for _, move := range moves {
		count += dfs(matrix, i+move.x, j+move.y, visited)
	}
	return count
}

func inBound(matrix [][]int, i, j int) bool {
	return i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0])
}
func Print(matrix [][]int) {
	for _, row := range matrix {
		for _, c := range row {
			if c == 0 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("# ")
			}
		}
		fmt.Println()
	}
}

type Bot struct {
	x, y, vx, vy int
}

func mod(x, y int) int {
	return (x%y + y) % y
}

func Matrix(h, w int) [][]int {
	var matrix [][]int
	for i := 0; i < h; i++ {
		var row []int
		for j := 0; j < w; j++ {
			row = append(row, 0)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
