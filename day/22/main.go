package main

import (
	"fmt"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")
	var secrets []int
	for _, line := range input {
		secrets = append(secrets, utils.ToInt(line))
	}

	total := 0
	for i := range secrets {
		for j := 0; j < 2000; j++ {
			secrets[i] = Transform(secrets[i])
		}
		total += secrets[i]
	}
	fmt.Println(total)
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
