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

	var matrix Matrix = utils.ToArr2D(input[:i])
	actions := strings.Join(input[i+1:], "")

	p := matrix.Start()
	matrix.Print()
	for _, action := range actions {
		m := Direction(action)
		i, j := p.x+m.x, p.y+m.y
		for matrix[i][j] == 'O' {
			i += m.x
			j += m.y
		}
		if matrix[i][j] == '#' {
			continue
		}
		matrix[i][j] = 'O'
		matrix[p.x][p.y] = '.'
		p.x += m.x
		p.y += m.y
		matrix[p.x][p.y] = '@'
	}
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
			} else {
				rr = append(rr, c)
				rr = append(rr, c)
			}
		}
		mm = append(mm, rr)
	}
	return mm
}
func (m Matrix) Cost() int {
	total := 0
	for i, row := range m {
		for j, c := range row {
			if c == 'O' {
				total += 100*i + j
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

func (m Matrix) Start() Move {
	for i, row := range m {
		for j, c := range row {
			if c == '@' {
				return Move{i, j}
			}
		}
	}
	panic("WTU?")
}
