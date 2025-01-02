package main

import (
	"fmt"
	"strconv"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	keypad := []string{"789", "456", "123", "X0A"}
	touchpad := []string{"X^A", "<v>"}
	total := 0
	for _, code := range input {
		n := len(code)
		plen := solve(keypad, touchpad, code)
		num, _ := strconv.Atoi(code[:n-1])
		total += num * plen
	}
	fmt.Println(total)
}

func solve(keypad, touchpad []string, code string) int {
	p1 := translate(keypad, code)

	p2 := make([]string, 0)
	for _, p := range p1 {
		t := translate(touchpad, p)
		p2 = append(p2, t...)
	}

	p3 := make([]string, 0)
	for _, p := range p2 {
		t := translate(touchpad, p)
		p3 = append(p3, t...)
	}

	mx := 1_000_000
	for _, p := range p3 {
		mx = min(mx, len(p))
	}
	return mx
}

func translate(pad []string, code string) []string {
	p := find(pad, 'A')
	paths := []string{""}
	for _, c := range code {
		paths = Product(paths, search(pad, p, find(pad, c)))
		p = find(pad, c)
	}
	return paths
}

func Product(p, q []string) []string {
	res := make([]string, 0, len(p)*len(q))
	for _, s := range p {
		for _, t := range q {
			res = append(res, s+t)
		}
	}
	return res
}

func search(pad []string, p, q Pos) []string {
	queue := []Item{{p, ""}}
	visited := make(map[Pos]bool)
	moves := []Move{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	paths := make([]string, 0)
	mx := 1_000_000_000
	for len(queue) > 0 {
		p, path := queue[0].p, queue[0].path
		queue = queue[1:]
		if !bound(pad, p) || visited[p] || pad[p.x][p.y] == 'X' {
			continue
		}
		if len(path) > mx {
			continue
		}
		if p.x == q.x && p.y == q.y {
			mx = min(len(path), mx)
			paths = append(paths, path+"A")
			continue
		}
		for _, m := range moves {
			pnew := Pos{p.x + m.dx, p.y + m.dy}
			queue = append(queue, Item{pnew, path + m.dir()})
		}
	}
	return paths
}

type Move struct {
	dx, dy int
}

func find(pad []string, s rune) Pos {
	for i := 0; i < len(pad); i++ {
		for j := 0; j < len(pad[0]); j++ {
			if pad[i][j] == byte(s) {
				return Pos{i, j}
			}
		}
	}
	return Pos{}
}

func (m Move) dir() string {
	switch {
	case m.dx == 1 && m.dy == 0:
		return "v"
	case m.dx == -1 && m.dy == 0:
		return "^"
	case m.dx == 0 && m.dy == 1:
		return ">"
	case m.dx == 0 && m.dy == -1:
		return "<"
	default:
		panic("???")
	}
}

type Item struct {
	p    Pos
	path string
}

type Pos struct {
	x, y int
}

func bound(matrix []string, p Pos) bool {
	return p.x >= 0 && p.x < len(matrix) && p.y >= 0 && p.y < len(matrix[0])
}
