package main

import (
	"fmt"
	"strings"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")

	var bots []Bot
	for _, line := range input {
		b := strings.Split(line, ",")
		bots = append(bots, Bot{
			utils.ToInt(b[1]),
			utils.ToInt(b[0]),
			utils.ToInt(b[3]),
			utils.ToInt(b[2]),
		})
	}
	t := 100
	h, w := 103, 101
	hm, wm := (h)/2, (w)/2
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, bot := range bots {
		px := mod(bot.x+t*bot.vx, h)
		py := mod(bot.y+t*bot.vy, w)

		if (px < hm) && (py < wm) {
			q1++
		}
		if (px < hm) && (py > wm) {
			q2++
		}
		if (px > hm) && (py < wm) {
			q3++
		}
		if (px > hm) && (py > wm) {
			q4++
		}
		fmt.Println(bot, px, py)
	}
	fmt.Println(q1, q2, q3, q4, q1*q2*q3*q4)
}

type Bot struct {
	x, y, vx, vy int
}

func mod(x, y int) int {
	return (x%y + y) % y
}
