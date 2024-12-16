package main

import (
	"fmt"

	"container/heap"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	matrix := utils.ToArr2D(input)

	p := Start(matrix)
	cost := search(matrix, p)
	fmt.Println(cost)
}

type Position struct {
	p Point
	d Direction
}

func search(matrix [][]rune, p Point) int {
	pq := make(PriorityQueue, 0)
	best := make(map[Position]int)
	paths := make([]Point, 0)
	max_cost := int(1e9)

	heap.Init(&pq)
	heap.Push(&pq, &Item{cost: 0, pos: p, dir: "E", path: []Point{}})
	for len(pq) > 0 {
		i := heap.Pop(&pq).(*Item)
		cost, p, d := i.cost, i.pos, i.dir
		path := Copy(i.path, p)

		if matrix[p.x][p.y] == '#' {
			continue
		}

		if cost > max_cost {
			continue
		}

		if c, ok := best[Position{p, d}]; ok && cost > c {
			continue
		}
		best[Position{p, d}] = cost

		if matrix[p.x][p.y] == 'E' {
			paths = append(paths, path...)
			max_cost = cost
		}

		i1 := Item{pos: p.Move(d), dir: d, cost: i.cost + 1, path: path}
		i2 := Item{pos: p.Move(d.TurnLeft()), dir: d.TurnLeft(), cost: i.cost + 1001, path: path}
		i3 := Item{pos: p.Move(d.TurnRight()), dir: d.TurnRight(), cost: i.cost + 1001, path: path}

		heap.Push(&pq, &i1)
		heap.Push(&pq, &i2)
		heap.Push(&pq, &i3)
	}

	seen := make(map[Point]bool)
	for _, p := range paths {
		seen[p] = true
	}
	fmt.Println(len(seen))
	return max_cost
}

func Copy(slice []Point, p Point) []Point {
	res := make([]Point, len(slice))
	copy(res, slice)
	return append(res, p)
}

type Point struct {
	x, y int
}

func (p Point) Move(d Direction) Point {
	return Point{p.x + d.Move().dx, p.y + d.Move().dy}
}

type Move struct {
	dx, dy int
}

type Direction string

func (d Direction) TurnLeft() Direction {
	switch d {
	case "N":
		return "W"
	case "W":
		return "S"
	case "S":
		return "E"
	case "E":
		return "N"
	default:
		panic("WUT?")
	}
}

func (d Direction) TurnRight() Direction {
	switch d {
	case "N":
		return "E"
	case "E":
		return "S"
	case "S":
		return "W"
	case "W":
		return "N"
	default:
		panic("WUT?")
	}
}

func (d Direction) Move() Move {
	switch d {
	case "N":
		return Move{-1, 0}
	case "E":
		return Move{0, 1}
	case "S":
		return Move{1, 0}
	case "W":
		return Move{0, -1}
	default:
		panic("WUT?")
	}
}

func Start(matrix [][]rune) Point {
	for i, row := range matrix {
		for j, c := range row {
			if c == 'S' {
				return Point{i, j}
			}
		}
	}
	panic("WTU?")
}
