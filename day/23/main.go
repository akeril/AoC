package main

import (
	"fmt"
	"strings"

	"github.com/kjabin/aoc2024/utils"
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

	queue := make([][]string, 0)
	for u := range adj {
		queue = append(queue, []string{u})
	}

	conns := make([][]string, 0)
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		u := path[len(path)-1]

		if len(path) > 4 {
			continue
		}

		if len(path) == 4 && path[0] == u {
			conns = append(conns, path)
			continue
		}

		for _, v := range adj[u] {
			queue = append(queue, Copy(path, v))
		}
	}

	count := 0
	for _, conn := range conns {
		for _, u := range conn {
			if u[0] == 't' {
				count++
				break
			}
		}
	}
	fmt.Println(count / 6)

}

func Copy(slice []string, p string) []string {
	res := make([]string, len(slice))
	copy(res, slice)
	return append(res, p)
}

type Conn struct {
	x, y, z string
}

func dfs(adj map[string][]string, visited map[string]bool, u string) {
	if visited[u] {
		return
	}
	visited[u] = true
	fmt.Printf("%s ", u)
	for _, v := range adj[u] {
		dfs(adj, visited, v)
	}
}
