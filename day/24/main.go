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

	// maintains values for each wire
	wires := make(map[string]int)
	for _, line := range input[:k] {
		var key string
		var value int
		fmt.Sscanf(line, "%3s: %d", &key, &value)
		wires[key] = value
	}

	// maintain formulas for each gate
	deps := make(map[string]Gate)
	// maintains calculation order
	graph := make(map[string][]string)
	for _, line := range input[k+1:] {
		var w1, w2, w3, op string
		fmt.Sscanf(line, "%3s %s %3s -> %3s", &w1, &op, &w2, &w3)
		graph[w1] = append(graph[w1], w3)
		graph[w2] = append(graph[w2], w3)
		deps[w3] = Gate{w1, w2, op, 2}
	}

	queue := make([]string, 0)
	for w := range wires {
		queue = append(queue, w)
	}

	for len(queue) > 0 {
		w := queue[0]
		queue = queue[1:]
		if _, ok := wires[w]; !ok {
			wires[w] = deps[w].Calc(wires)
		}
		for _, v := range graph[w] {
			g := deps[v]
			g.in--
			deps[v] = g

			if g.in == 0 {
				queue = append(queue, v)
			}
		}
	}

	zwires := make([]Wire, 0)
	for wire, value := range wires {
		if wire[0] == 'z' {
			zwires = append(zwires, Wire{wire, value})
		}
	}

	slices.SortFunc(zwires, func(a, b Wire) int {
		return -cmp.Compare(a.id, b.id)
	})

	bits := ""
	for _, wire := range zwires {
		bits += strconv.Itoa(wire.value)
	}
	fmt.Println(strconv.ParseInt(bits, 2, 64))
}

type Wire struct {
	id    string
	value int
}

type Gate struct {
	w1, w2 string
	op     string
	in     int
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
