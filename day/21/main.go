package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/akeril/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	keypad := []string{"789", "456", "123", "X0A"}
	directionalpad := []string{"X^A", "<v>"}
	directionalseq := seq(directionalpad)
	cache := make(map[citem]int)
	total := 0
	for _, code := range input {
		mi := math.MaxInt
		for _, path := range search(keypad, code) {
			c := 0
			for _, p := range zip("A" + path) {
				c += cost(cache, directionalseq, p.s, p.e, 25)
			}
			mi = min(mi, c)
		}
		num, _ := strconv.Atoi(code[:len(code)-1])
		total += num * mi
	}
	fmt.Println(total)
}

type citem struct {
	p pair
	d int
}

type pair struct {
	s, e rune
}

func cost(cache map[citem]int, seq map[pair][]string, start, end rune, depth int) int {
	p := pair{start, end}
	ci := citem{p, depth}
	if depth == 1 {
		return minlenstr(seq[p])
	}
	if v, ok := cache[ci]; ok {
		return v
	}
	mi := math.MaxInt
	for _, path := range seq[p] {
		c := 0
		for _, pp := range zip("A" + path) {
			c += cost(cache, seq, pp.s, pp.e, depth-1)
		}
		mi = min(mi, c)
	}
	cache[ci] = mi
	return mi
}

func zip(s string) []pair {
	var res []pair
	for i := 1; i < len(s); i++ {
		res = append(res, pair{rune(s[i-1]), rune(s[i])})
	}
	return res
}

func seq(pad []string) map[pair][]string {
	mp := make(map[pair][]string)
	s := strings.Join(pad, "")
	for _, start := range s {
		for _, end := range s {
			mp[pair{start, end}] = bfs(pad, start, end)
		}
	}
	return mp
}

func search(pad []string, code string) []string {
	paths := []string{""}
	for _, p := range zip("A" + code) {
		paths = product(paths, bfs(pad, p.s, p.e))
	}
	return paths
}

func product(ss []string, tt []string) []string {
	var res []string
	for _, s := range ss {
		for _, t := range tt {
			res = append(res, s+t)
		}
	}
	return res
}

// find all shortest path between two points
func bfs(pad []string, start, end rune) []string {
	if start == 'X' || end == 'X' {
		return []string{}
	}
	type item struct {
		p    pos
		path string
	}
	maxlen := math.MaxInt
	paths := make([]string, 0)
	queue := []item{{find(pad, start), ""}}
	moves := []move{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for len(queue) > 0 {
		p, path := queue[0].p, queue[0].path
		queue = queue[1:]
		if !bound(pad, p) || pad[p.x][p.y] == 'X' {
			continue
		}
		if len(path) > maxlen {
			continue
		}
		if pad[p.x][p.y] == byte(end) {
			path = path + "A"
			maxlen = len(path)
			paths = append(paths, path)
			continue
		}
		for _, m := range moves {
			it := item{pos{p.x + m.dx, p.y + m.dy}, path + m.dir()}
			queue = append(queue, it)
		}
	}
	return paths
}

type move struct {
	dx, dy int
}

func (m move) dir() string {
	switch {
	case m.dx == 1 && m.dy == 0:
		return "v"
	case m.dx == -1 && m.dy == 0:
		return "^"
	case m.dx == 0 && m.dy == 1:
		return ">"
	case m.dx == 0 && m.dy == -1:
		return "<"
	}
	return "?"
}

type pos struct {
	x, y int
}

func bound(pad []string, p pos) bool {
	return p.x >= 0 && p.x < len(pad) && p.y >= 0 && p.y < len(pad[0])
}

func find(pad []string, c rune) pos {
	for i, row := range pad {
		for j, cc := range row {
			if c == cc {
				return pos{i, j}
			}
		}
	}
	return pos{}
}

func minlenstr(ss []string) int {
	mi := math.MaxInt
	for _, s := range ss {
		mi = min(mi, len(s))
	}
	return mi
}
