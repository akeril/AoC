package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/akeril/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	adj := make(map[string][]string)
	for _, line := range input {
		nodes := strings.Split(line, "-")
		u, v := nodes[0], nodes[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	for u := range adj {
		visited := make(map[string]bool)
		components := make([]string, 0)
		dfs(adj, visited, &components, u)
		fmt.Println(strings.Join(components, ","))
	}
}

func Copy(slice []string, p string) []string {
	res := make([]string, len(slice))
	copy(res, slice)
	return append(res, p)
}

type Conn struct {
	x, y, z string
}

func dfs(adj map[string][]string, visited map[string]bool, components *[]string, u string) {
	if visited[u] {
		return
	}
	visited[u] = true
	*components = append(*components, u)
	for _, v := range adj[u] {
		if Subset(adj[v], Keys(visited)) {
			dfs(adj, visited, components, v)
		}
	}
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Subset[K comparable](l, s []K) bool {
	for _, x := range s {
		if !slices.Contains(l, x) {
			return false
		}
	}
	return true
}
