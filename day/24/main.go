package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	k := slices.Index(input, "")

	values := make(map[string]int)
	for _, line := range input[:k] {
		var key string
		var value int
		fmt.Sscanf(line, "%3s: %d", &key, &value)
		values[key] = value
	}

	formulas := make(map[string]Formula)
	for _, line := range input[k+1:] {
		var w1, w2, w3, op string
		fmt.Sscanf(line, "%3s %s %3s -> %3s", &w1, &op, &w2, &w3)
		formulas[w3] = Formula{w1, w2, op}
	}

	for w := range formulas {
		values[w] = calc(values, formulas, w)
	}

	fmt.Println(reg(values, 'z'))
}

func reg(values map[string]int, c byte) int64 {
	i := 0
	bits := ""
	for {
		reg := fmt.Sprintf("%c%02d", c, i)
		if _, ok := values[reg]; !ok {
			break
		}
		bits = strconv.Itoa(values[reg]) + bits
		i += 1
	}
	num, _ := strconv.ParseInt(bits, 2, 64)
	return num
}

func calc(values map[string]int, formulas map[string]Formula, w string) int {
	if v, ok := values[w]; ok {
		return v
	}
	f := formulas[w]
	values[w] = operate(f.op, calc(values, formulas, f.w1), calc(values, formulas, f.w2))
	return values[w]
}

func operate(op string, w1, w2 int) int {
	switch op {
	case "AND":
		return w1 & w2
	case "OR":
		return w1 | w2
	case "XOR":
		return w1 ^ w2
	default:
		return 0
	}
}

type Formula struct {
	w1, w2 string
	op     string
}
