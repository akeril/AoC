package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	i := slices.Index(input, "")

	matrix := Matrix(utils.ToArr2D(input[:i]))
	actions := strings.Join(input[i+1:], "")

	matrix = matrix.Double()
	p := matrix.StartPos()
	for _, action := range actions {
		d := Direction(action)
		if matrix.TryPush(p, d) {
			matrix.Push(p, d)
			p.x += d.x
			p.y += d.y
		}
	}
	matrix.Print()
	fmt.Println(matrix.Cost())
}

type Matrix [][]rune

func (m Matrix) Double() Matrix {
	var mm [][]rune
	for _, row := range m {
		var rr []rune
		for _, c := range row {
			if c == 'O' {
				rr = append(rr, '[')
				rr = append(rr, ']')
			} else if c == '@' {
				rr = append(rr, '@')
				rr = append(rr, '.')
			} else {
				rr = append(rr, c)
				rr = append(rr, c)
			}
		}
		mm = append(mm, rr)
	}
	return mm
}

func (m Matrix) Push(p, d Move) {
	if m[p.x][p.y] == '.' {
		return
	}
	if m[p.x][p.y] == '#' {
		return
	}
	if m[p.x][p.y] == 'V' {
		return
	}

	c := m[p.x][p.y]
	m[p.x][p.y] = 'V'
	if c == '@' {
		m.Push(Move{p.x + d.x, p.y + d.y}, d)
	} else if d.y == 0 && c == '[' {
		m.Push(Move{p.x, p.y + 1}, d)
		m.Push(Move{p.x + d.x, p.y + d.y}, d)
	} else if d.y == 0 && c == ']' {
		m.Push(Move{p.x, p.y - 1}, d)
		m.Push(Move{p.x + d.x, p.y + d.y}, d)
	} else {
		m.Push(Move{p.x + d.x, p.y + d.y}, d)
	}

	m[p.x+d.x][p.y+d.y] = c
	m[p.x][p.y] = '.'
}

func (m Matrix) TryPush(p, d Move) bool {
	if m[p.x][p.y] == '.' {
		return true
	}

	if m[p.x][p.y] == '#' {
		return false
	}

	if m[p.x][p.y] == '@' {
		return m.TryPush(Move{p.x + d.x, p.y + d.y}, d)
	}

	// vertical movement
	if d.y == 0 && m[p.x][p.y] == '[' {
		return m.TryPush(Move{p.x + d.x, p.y + d.y}, d) && m.TryPush(Move{p.x + d.x, p.y + d.y + 1}, d)
	}

	if d.y == 0 && m[p.x][p.y] == ']' {
		return m.TryPush(Move{p.x + d.x, p.y + d.y}, d) && m.TryPush(Move{p.x + d.x, p.y + d.y - 1}, d)
	}

	return m.TryPush(Move{p.x + d.x, p.y + d.y}, d)
}

func (m Matrix) Cost() int {
	total := 0
	for i, row := range m {
		for j, c := range row {
			if c == '[' {
				total += 100*i + (j)
			}
		}
	}
	return total
}

func (m Matrix) Print() {
	for _, row := range m {
		for _, c := range row {
			fmt.Printf(" %v", string(c))
		}
		fmt.Println()
	}
}

func (m Matrix) inBound(i, j int) bool {
	return i >= 0 && i < len(m) && j >= 0 && j < len(m[0])
}

type Move struct {
	x, y int
}

func Direction(r rune) Move {
	switch r {
	case '^':
		return Move{-1, 0}
	case 'v':
		return Move{1, 0}
	case '<':
		return Move{0, -1}
	case '>':
		return Move{0, 1}
	default:
		panic(string(r) + " WUT?")
	}
}

func (m Matrix) StartPos() Move {
	for i, row := range m {
		for j, c := range row {
			if c == '@' {
				return Move{i, j}
			}
		}
	}
	panic("WTU?")
}
