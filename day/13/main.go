package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input, _ := os.ReadFile("input")

	puzzles := strings.Split(string(input), "\n\n")

	count := 0
	for _, puzzle := range puzzles {
		buttons := strings.Split(puzzle, "\n")
		a := ParseValues(buttons[0])
		b := ParseValues(buttons[1])
		c := ParseValues(buttons[2])
		// ax.i + bx.j = cx
		// ay.i + by.j = cy
		// ax.ay.i + ay.bx.j = ay.cx
		// ax.ay.i + ax.by.j = ax.cy
		// j = (ay.cx - ax.cy) / (ay.bx - ax.by)
		// i = (cx - bx.j) / ax

		c.x += 10000000000000
		c.y += 10000000000000
		j := (a.y*c.x - a.x*c.y) / (a.y*b.x - a.x*b.y)
		i := (c.x - b.x*j) / a.x

		if i < 0 || j < 0 {
			continue
		}
		if a.x*i+b.x*j != c.x {
			continue
		}
		if a.y*i+b.y*j != c.y {
			continue
		}
		fmt.Println(a, b, c)
		count += 3*i + j
	}
	fmt.Println(count)
}

type Vector struct {
	x, y int
}

func ParseValues(input string) Vector {
	inputs := strings.Split(input, ",")
	x := utils.ToInt(ExtractInt(inputs[0]))
	y := utils.ToInt(ExtractInt(inputs[1]))
	return Vector{x, y}
}

func ExtractInt(x string) string {
	s := ""
	for _, c := range x {
		if c >= '0' && c <= '9' {
			s += string(c)
		}
	}
	return s
}
