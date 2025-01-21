package main

import (
	"fmt"

	"github.com/akeril/aoc2024/utils"
)

type pos struct {
	x, y int
}

func main() {
	input := utils.ReadFile("input")
	fmt.Println(calcAntiNodes(input))
}

func calcAntiNodes(matrix []string) int {
	nodes := make(map[rune][]pos)
	for i, line := range matrix {
		for j, c := range line {
			if c != '.' {
				nodes[c] = append(nodes[c], pos{i, j})
			}
		}
	}
	antinodes := make(map[pos]bool)
	for _, points := range nodes {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p, q := points[i], points[j]
				for k := -100; k < 100; k++ {
					a := pos{k*(p.x-q.x) + q.x, k*(p.y-q.y) + q.y}
					if a.inBound(matrix) {
						antinodes[a] = true
					}
				}
			}
		}
	}
	return len(antinodes)
}

// x + y = 2z - y
func (p pos) inBound(matrix []string) bool {
	i, j := p.x, p.y
	return i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0])
}
