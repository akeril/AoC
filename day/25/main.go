package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input")
	schemas := strings.Split(string(input), "\n\n")
	locks, keys := make([]Schema, 0), make([]Schema, 0)

	for _, schema := range schemas {
		pattern := strings.Split(strings.Trim(schema, "\n"), "\n")
		n, m := len(pattern[0]), len(pattern)
		counts := make([]int, n)
		for i := 0; i < n; i++ {
			j := 0
			for j < m && pattern[j][i] == pattern[0][i] {
				j++
			}
			if pattern[0][0] == '#' {
				counts[i] = j
			} else {
				counts[i] = m - j
			}
		}
		sch := Schema{counts: counts, n: n, m: m}
		if pattern[0][0] == '#' {
			locks = append(locks, sch)
		} else {
			keys = append(keys, sch)
		}
	}

	total := 0
	for _, lock := range locks {
		for _, key := range keys {
			flag := true
			for i := 0; i < lock.n; i++ {
				if lock.counts[i]+key.counts[i] > lock.m {
					flag = false
				}
			}
			if flag {
				total++
			}
		}
	}
	fmt.Println(total)
}

type Schema struct {
	n, m   int
	counts []int
}
