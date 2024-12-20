package main

import (
	"fmt"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	matrix := utils.ToArr2D(input)
	FindShortcuts(matrix)
}

func FindShortcuts(matrix [][]rune) {
	costs := make(map[Point]int)
	Walk(costs, matrix, Find(matrix, 'S'), Point{})

	count := 0
	n := 20
	for s := range costs {
		for _, pos := range LocatePoints(matrix, s, n) {
			e, d := pos.Point, pos.d
			if costs[s]-costs[e]-d >= 100 {
				count++
			}
		}
	}
	fmt.Println(count)
}

type Cheat struct {
	s, e Point
}

func LocatePoints(matrix [][]rune, p Point, n int) []Position {
	queue := []Position{{p, 0}}
	result := make([]Position, 0)
	moves := []Move{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}
	visited := make(map[Point]bool)
	for len(queue) > 0 {
		u, d := queue[0].Point, queue[0].d
		queue = queue[1:]
		if d > n || !Bound(matrix, u.x, u.y) {
			continue
		}
		if visited[u] {
			continue
		}
		visited[u] = true
		if matrix[u.x][u.y] == '.' || matrix[u.x][u.y] == 'E' {
			result = append(result, Position{u, d})
		}
		for _, m := range moves {
			queue = append(queue, Position{Point{u.x + m.dx, u.y + m.dy}, d + 1})
		}
	}
	return result
}

func Walk(visited map[Point]int, matrix [][]rune, c, p Point) {
	if matrix[c.x][c.y] == 'E' {
		visited[c] = 0
	} else {
		n := Adj(matrix, c, p)
		Walk(visited, matrix, n, c)
		visited[c] = visited[n] + 1
	}
}

func Adj(matrix [][]rune, c, p Point) Point {
	moves := []Move{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}
	for _, m := range moves {
		n := Point{c.x + m.dx, c.y + m.dy}
		if (matrix[n.x][n.y] == 'E' || matrix[n.x][n.y] == '.') && n != p {
			return n
		}
	}
	panic("??")
}

type Position struct {
	Point
	d int
}

type Point struct {
	x, y int
}

type Move struct {
	dx, dy int
}

func Find(matrix [][]rune, s rune) Point {
	for i, row := range matrix {
		for j, c := range row {
			if c == s {
				return Point{i, j}
			}
		}
	}
	panic("WTU?")
}

func Bound(matrix [][]rune, i, j int) bool {
	return i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0])
}
