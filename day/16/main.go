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
	fmt.Println(search(matrix, p))
}

func search(matrix [][]rune, p Pos) int {
	pq := make(PriorityQueue, 0)
	visited := make(map[Pos]bool)

	cost := 99999999
	heap.Init(&pq)
	heap.Push(&pq, &Item{cost: 0, pos: p})
	for len(pq) > 0 {
		i := heap.Pop(&pq).(*Item)
		p := i.pos

		if matrix[p.x][p.y] == '#' {
			continue
		}

		if _, ok := visited[p]; ok {
			continue
		}
		visited[p] = true

		if i.cost > cost {
			continue
		}

		if matrix[p.x][p.y] == 'E' {
			cost = min(cost, i.cost)
			continue
		}
		i1 := Item{pos: p.Move(p.d), cost: i.cost + 1}
		i2 := Item{pos: p.Move(p.d.TurnLeft()), cost: i.cost + 1001}
		i3 := Item{pos: p.Move(p.d.TurnRight()), cost: i.cost + 1001}

		heap.Push(&pq, &i1)
		heap.Push(&pq, &i2)
		heap.Push(&pq, &i3)
	}

	return cost
}

type Pos struct {
	x, y int
	d    Direction
}

func (p Pos) Move(d Direction) Pos {
	p.d = d
	return Pos{p.x + p.d.Move().dx, p.y + p.d.Move().dy, p.d}
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

func Start(matrix [][]rune) Pos {
	for i, row := range matrix {
		for j, c := range row {
			if c == 'S' {
				return Pos{i, j, "E"}
			}
		}
	}
	panic("WTU?")
}
