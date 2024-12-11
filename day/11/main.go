package main

import (
	"fmt"
	"strconv"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	mp := make(map[int]int)
	for _, stone := range utils.ToIntArr(input[0], " ") {
		mp[stone] = 1
	}

	for i := 0; i < 75; i++ {
		m := make(map[int]int)
		for stone, count := range mp {
			for _, s := range Transform(stone) {
				m[s] += count
			}
		}
		mp = m
	}
	count := 0
	for _, v := range mp {
		count += v
	}
	fmt.Println(count)
}

func Transform(x int) []int {
	if x == 0 {
		return []int{1}
	}
	sx := strconv.Itoa(x)
	l := len(sx)
	if l%2 == 0 {
		return []int{utils.ToInt(sx[l/2:]), utils.ToInt(sx[:l/2])}
	}
	return []int{x * 2024}
}
