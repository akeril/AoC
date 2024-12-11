package main

import (
	"fmt"
	"strconv"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	stk1 := New()
	stk2 := New()
	PushArr(stk1, utils.ToIntArr(input[0], " "))

	stk1.Print()
	for i := 0; i < 75; i++ {
		fmt.Println(i)
		stk1, stk2 = stk2, stk1
		for stk2.Len() != 0 {
			v := stk2.Pop().(int)
			PushArr(stk1, Transform(v))
		}
	}
	fmt.Println(stk1.Len())
}

func PushArr(stk *Stack, arr []int) {
	for _, v := range arr {
		stk.Push(v)
	}
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
