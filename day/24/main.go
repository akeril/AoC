package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/akeril/aoc2024/utils"
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

	formulas["z11"], formulas["wpd"] = formulas["wpd"], formulas["z11"]
	formulas["jqf"], formulas["skh"] = formulas["skh"], formulas["jqf"]
	formulas["z19"], formulas["mdd"] = formulas["mdd"], formulas["z19"]
	formulas["z37"], formulas["wts"] = formulas["wts"], formulas["z37"]

	for w := range formulas {
		values[w] = calc(values, formulas, w)
	}

	zbits := reg(values, 'z')
	xbits := reg(values, 'x')
	ybits := reg(values, 'y')

	fmt.Printf("0%s\n", xbits)
	fmt.Printf("0%s\n", ybits)
	fmt.Println(zbits)
}

func pprint(values map[string]int, formulas map[string]Formula, w string, depth int) {
	if depth > 4 {
		return
	}
	if _, ok := formulas[w]; ok {
		f := formulas[w]
		fmt.Println(strings.Repeat(" ", depth), f.op, w)
		pprint(values, formulas, f.w1, depth+1)
		pprint(values, formulas, f.w2, depth+1)
	} else {
		fmt.Println(strings.Repeat(" ", depth), w)
	}
}

func reg(values map[string]int, c byte) string {
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
	return bits
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
