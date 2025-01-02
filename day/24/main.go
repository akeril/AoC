package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	k := slices.Index(input, "")

	wires := make(map[string]int)
	for _, line := range input[:k] {
		var key string
		var value int
		fmt.Sscanf(line, "%3s: %d", &key, &value)
		wires[key] = value
	}

	deps := make(map[string]Gate)
	for _, line := range input[k+1:] {
		var w1, w2, w3, op string
		fmt.Sscanf(line, "%3s %s %3s -> %3s", &w1, &op, &w2, &w3)
		deps[w3] = Gate{w1, w2, op}
	}

	for w := range deps {
		dfs(deps, wires, w)
	}

	fmt.Println(regVal(wires, 'z'))
}

func regVal(wires map[string]int, c byte) int64 {
	vals := make([]Wire, 0)
	for w, v := range wires {
		if w[0] == c {
			vals = append(vals, Wire{w, v})
		}
	}
	slices.SortFunc(vals, func(a, b Wire) int { return -cmp.Compare(a.id, b.id) })
	bits := ""
	for _, v := range vals {
		bits += strconv.Itoa(v.value)
	}
	num, _ := strconv.ParseInt(bits, 2, 64)
	return num
}

func dfs(deps map[string]Gate, wires map[string]int, w string) int {
	if _, ok := wires[w]; !ok {
		dfs(deps, wires, deps[w].w1)
		dfs(deps, wires, deps[w].w2)
		wires[w] = deps[w].Calc(wires)
	}
	return wires[w]
}

type Wire struct {
	id    string
	value int
}

type Gate struct {
	w1, w2 string
	op     string
}

func (g Gate) Calc(wires map[string]int) int {
	switch g.op {
	case "AND":
		return wires[g.w1] & wires[g.w2]
	case "OR":
		return wires[g.w1] | wires[g.w2]
	case "XOR":
		return wires[g.w1] ^ wires[g.w2]
	default:
		return 0
	}
}
