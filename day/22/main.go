package main

import (
	"fmt"
	"strconv"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	var secrets []int
	for _, line := range input {
		secrets = append(secrets, utils.ToInt(line))
	}

	adj := make(map[string][]Item)
	for i := range secrets {
		price := secrets[i] % 10
		deltas := make([]int, 4)
		for j := 0; j < 2000; j++ {
			secrets[i] = Transform(secrets[i])
			deltas[j%4] = (secrets[i] % 10) - price
			price = secrets[i] % 10
			if j >= 3 {
				hash := Hash(deltas, j)
				adj[hash] = append(adj[hash], Item{i, j, price})
			}
		}
	}

	mx := -1
	for _, prices := range adj {
		total := 0
		j := 0
		for i := range secrets {
			if j < len(prices) && prices[j].i == i {
				total += prices[j].price
			}
			for j < len(prices) && prices[j].i <= i {
				j++
			}
		}
		mx = max(total, mx)
	}
	fmt.Println(mx)
}

type Item struct {
	i, j, price int
}

func Hash(deltas []int, j int) string {
	res := ""
	for i := j%4 + 1; i < 4; i++ {
		res += strconv.Itoa(deltas[i])
	}
	for i := 0; i < j%4+1; i++ {
		res += strconv.Itoa(deltas[i])
	}
	return res
}

func Transform(secret int) int {
	secret = secret ^ (secret * 64)
	secret = secret % 16777216

	secret = secret ^ (secret / 32)
	secret = secret % 16777216

	secret = secret ^ (secret * 2048)
	secret = secret % 16777216
	return secret
}
