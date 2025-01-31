package utils

import (
	"strconv"
	"strings"
)

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func ToIntArr(s string, sep string) []int {
	var result []int
	for _, num := range strings.Split(s, sep) {
		result = append(result, ToInt(num))
	}
	return result
}

func ToArr2D(ss []string) [][]rune {
	var result [][]rune
	for _, s := range ss {
		result = append(result, []rune(s))
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
