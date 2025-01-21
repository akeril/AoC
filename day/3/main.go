package main

import (
	"fmt"
	"regexp"
	"strings"

	utl "github.com/akeril/aoc2024/utils"
)

func main() {
	lines := utl.ReadFile("input")
	query := strings.Join(lines, "")
	fmt.Println(calcQuery(query))
	fmt.Println(calcQueryII(query))
}

func calcQuery(query string) int {
	total := 0
	re := regexp.MustCompile(`mul\((-?\d+),(-?\d+)\)`)
	for _, match := range re.FindAllStringSubmatch(query, -1) {
		total += utl.ToInt(match[1]) * utl.ToInt(match[2])
	}
	return total
}

func calcQueryII(query string) int {
	total := 0
	enabled := true
	re := regexp.MustCompile(`(do\(\))|(don't\(\))|(mul\((-?\d+),(-?\d+)\))`)
	for _, match := range re.FindAllStringSubmatch(query, -1) {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			total += utl.ToInt(match[4]) * utl.ToInt(match[5])
		}
	}
	return total
}
