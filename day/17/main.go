package main

import (
	"fmt"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	r := Register{
		A: utils.ToInt(input[0]),
		B: utils.ToInt(input[1]),
		C: utils.ToInt(input[2]),
	}
	instructions := utils.ToIntArr(input[3], ",")
	Solve(instructions, r)
}

func Solve(instructions []int, r Register) {
	i := 0
	output := make([]int, 0)
	for i < len(instructions) {
		opcode, operand := instructions[i], instructions[i+1]
		jmp := r.Command(opcode, operand, &output)
		if jmp < 0 {
			i += 2
		} else {
			i = jmp
		}
	}
	fmt.Println(output)
}

type Register struct {
	A, B, C int
}

func Pow(x, y int) int {
	res := 1
	for i := 0; i < y; i++ {
		res *= x
	}
	return res
}

func (r *Register) Command(opcode, operand int, output *[]int) int {
	res := -1
	switch opcode {
	case 0:
		r.A /= Pow(2, r.Combo(operand))
	case 1:
		r.B ^= operand
	case 2:
		r.B = r.Combo(operand) % 8
	case 3:
		if r.A != 0 {
			res = operand
		}
	case 4:
		r.B ^= r.C
	case 5:
		*output = append(*output, r.Combo(operand)%8)
		res = -2
	case 6:
		r.B = r.A / Pow(2, r.Combo(operand))
	case 7:
		r.C = r.A / Pow(2, r.Combo(operand))
	}
	return res
}

func (r Register) Combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return r.A
	case 5:
		return r.B
	case 6:
		return r.C
	default:
		panic("???")
	}
}
