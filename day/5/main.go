package main

import (
	"fmt"
	"slices"
	"strings"

	utl "github.com/akeril/aoc2024/utils"
)

func main() {
	input := utl.ReadFile("input")
	marker := slices.Index(input, "")
	graph, updates := createGraph(input[:marker]), createQueries(input[marker+1:])
	fmt.Println(totalCost(graph, updates))
	fmt.Println(totalCostII(graph, updates))
}

func totalCostII(graph map[int][]int, updates [][]int) int {
	total := 0
	for _, update := range updates {
		if !validUpdate(graph, update) {
			updated := findValidUpdate(graph, update)
			total += updated[len(updated)/2]
		}
	}
	return total
}

func totalCost(graph map[int][]int, updates [][]int) int {
	total := 0
	for _, update := range updates {
		if validUpdate(graph, update) {
			total += update[len(update)/2]
		}
	}
	return total
}

func findValidUpdate(graph map[int][]int, update []int) []int {
	fan := make(map[int]int)
	visited := make(map[int]bool)
	var queue []int
	for _, page := range update {
		fan[page] = 0
		visited[page] = false
	}
	for _, page := range update {
		for _, dep := range graph[page] {
			fan[dep]++
		}
	}
	for _, page := range update {
		if fan[page] == 0 {
			queue = append(queue, page)
		}
	}
	var updated []int
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		visited[u] = true
		updated = append(updated, u)

		for _, v := range graph[u] {
			if slices.Contains(update, v) && !visited[v] {
				fan[v]--
				if fan[v] == 0 {
					queue = append(queue, v)
				}
			}
		}

	}
	return updated
}

func validUpdate(graph map[int][]int, update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		u, v := update[i], update[i+1]
		if !slices.Contains(graph[u], v) {
			return false
		}
	}
	return true
}

func createGraph(orders []string) map[int][]int {
	graph := make(map[int][]int)
	for _, order := range orders {
		uv := strings.Split(order, "|")
		u, v := utl.ToInt(uv[0]), utl.ToInt(uv[1])
		graph[u] = append(graph[u], v)
	}
	return graph
}

func createQueries(updates []string) [][]int {
	var queries [][]int
	for _, update := range updates {
		query := utl.ToIntArr(update, ",")
		queries = append(queries, query)
	}
	return queries
}
