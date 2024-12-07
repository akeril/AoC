package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	results, operands := ParseInput(input)
	fmt.Println(validOperations(results, operands))
}

func validOperations(results []int, operands [][]int) int {
	total := 0
	for i := 0; i < len(results); i++ {
		if search(results[i], operands[i][1:], operands[i][0]) {
			total += results[i]
		}
	}
	return total
}

func search(result int, operands []int, curr int) bool {
	if len(operands) == 0 && curr == result {
		return true
	}
	if len(operands) == 0 {
		return false
	}
	b1 := search(result, operands[1:], curr*operands[0])
	b2 := search(result, operands[1:], curr+operands[0])
	concat := utils.ToInt(strconv.Itoa(curr) + strconv.Itoa(operands[0]))
	b3 := search(result, operands[1:], concat)
	return b1 || b2 || b3
}

func ParseInput(input []string) ([]int, [][]int) {
	var results []int
	var operands [][]int
	for _, line := range input {
		query := strings.Split(line, ":")
		ops := strings.Trim(query[1], " ")
		results = append(results, utils.ToInt(query[0]))
		operands = append(operands, utils.ToIntArr(ops, " "))
	}
	return results, operands
}
