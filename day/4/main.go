package main

import (
	"fmt"

	utl "github.com/akeril/aoc2024/utils"
)

func main() {
	matrix := utl.ReadFile("input")
	fmt.Println(countXMAS(matrix))
	fmt.Println(countMAS(matrix))
}

func countMAS(matrix []string) int {
	count := 0
	for i, line := range matrix {
		for j, c := range line {
			if c == 'A' && inBox(matrix, i, j) {
				l := matrix[i-1][j-1] + matrix[i+1][j+1]
				r := matrix[i+1][j-1] + matrix[i-1][j+1]
				if l == ('S'+'M') && r == ('S'+'M') {
					count += 1
				}
			}
		}
	}
	return count
}

func countXMAS(matrix []string) int {
	count := 0
	for i, line := range matrix {
		for j, c := range line {
			if c == 'X' {
				count += countWord(matrix, "XMAS", i, j)
			}
		}
	}
	return count
}

func countWord(matrix []string, word string, i, j int) int {
	moves := []int{-1, 0, 1}
	matchedWords := 0
	for _, x := range moves {
		for _, y := range moves {
			curr := ""
			for k := 0; k < 4; k++ {
				if inBound(matrix, i+k*x, j+k*y) {
					curr += string(matrix[i+k*x][j+k*y])
				}
			}
			if curr == word {
				matchedWords += 1
			}
		}
	}
	return matchedWords
}

func inBound(matrix []string, i, j int) bool {
	return i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0])
}

func inBox(matrix []string, i, j int) bool {
	return inBound(matrix, i-1, j-1) && inBound(matrix, i+1, j+1)
}
