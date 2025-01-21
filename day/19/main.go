package main

import (
	"fmt"
	"strings"

	"github.com/akeril/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	patterns := strings.Split(input[0], ", ")

	trie := New()
	for _, pattern := range patterns {
		trie.Insert(pattern)
	}

	dp := make(map[string]Value)
	count := 0
	valid := 0
	for _, design := range input[2:] {
		value := solve(dp, trie, design)
		if value.b {
			valid += 1
		}
		count += value.v
	}
	fmt.Println(valid, count)
}

type Value struct {
	b bool
	v int
}

func solve(dp map[string]Value, trie *Trie, design string) Value {
	if design == "" {
		return Value{true, 1}
	}
	if _, ok := dp[design]; ok {
		return dp[design]
	}
	value := Value{false, 0}
	for i := 1; i <= len(design); i++ {
		if trie.Contains(design[:i]) {
			partial := solve(dp, trie, design[i:])
			value.b = value.b || partial.b
			value.v = value.v + partial.v
		}
	}
	dp[design] = value
	return dp[design]
}

type Trie struct {
	valid    bool
	children []*Trie
}

func New() *Trie {
	return &Trie{
		valid:    false,
		children: make([]*Trie, 26),
	}
}

func (t *Trie) Insert(s string) {
	for _, c := range s {
		c = c - 'a'
		if t.children[c] == nil {
			t.children[c] = New()
		}
		t = t.children[c]
	}
	t.valid = true
}

func (t *Trie) Contains(s string) bool {
	for _, c := range s {
		c = c - 'a'
		if t.children[c] == nil {
			return false
		}
		t = t.children[c]
	}
	return t.valid
}
