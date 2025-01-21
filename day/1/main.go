package main

import (
	"fmt"
	"slices"
	"strings"

	utl "github.com/akeril/aoc2024/utils"
)

func main() {
	lines := utl.ReadFile("input")
	lloc, rloc := ParseInput(lines)
	fmt.Println(CalcDiff(lloc, rloc))
	fmt.Println(CalcSim(lloc, rloc))
}

func ParseInput(lines []string) ([]int, []int) {
	var lloc, rloc []int
	for _, line := range lines {
		nums := strings.Split(line, "   ")
		l, r := utl.ToInt(nums[0]), utl.ToInt(nums[1])
		lloc = append(lloc, l)
		rloc = append(rloc, r)
	}
	return lloc, rloc
}

func CalcDiff(lloc, rloc []int) int {
	slices.Sort(lloc)
	slices.Sort(rloc)
	diff := 0
	for i := 0; i < len(lloc); i++ {
		diff += utl.Abs(lloc[i] - rloc[i])
	}
	return diff
}

func CalcSim(lloc, rloc []int) int {
	counter := make(map[int]int)
	for _, num := range rloc {
		counter[num] += 1
	}
	sim := 0
	for _, num := range lloc {
		sim += num * counter[num]
	}
	return sim
}
