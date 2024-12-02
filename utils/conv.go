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

func ToIntArr(s string) []int {
	var result []int
	for _, num := range strings.Split(s, " ") {
		result = append(result, ToInt(num))
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
