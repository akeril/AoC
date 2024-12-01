package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Input(file string) ([]int, []int) {
	f, _ := os.Open(file)
	sc := bufio.NewScanner(f)

	l1 := make([]int, 0)
	l2 := make([]int, 0)
	for sc.Scan() {
		locations := strings.Split(sc.Text(), "   ")
		v1, _ := strconv.Atoi(locations[0])
		v2, _ := strconv.Atoi(locations[1])
		l1 = append(l1, v1)
		l2 = append(l2, v2)
	}
	return l1, l2
}

func Part1(l1, l2 []int) int {
	sort.Slice(l1, func(i, j int) bool { return l1[i] < l1[j] })
	sort.Slice(l2, func(i, j int) bool { return l2[i] < l2[j] })

	diff := 0
	for i := 0; i < len(l1); i++ {
		if l1[i] < l2[i] {
			diff += l2[i] - l1[i]
		} else {
			diff += l1[i] - l2[i]
		}
	}
	return diff
}

func Part2(l1, l2 []int) int {
	counter := make(map[int]int)
	for _, x := range l2 {
		counter[x] += 1
	}
	sim := 0
	for _, x := range l1 {
		if y, ok := counter[x]; ok {
			sim += x * y
		}
	}
	return sim
}

func main() {
	l1, l2 := Input("input")
	// fmt.Println(Part1(l1, l2))
	fmt.Println(Part2(l1, l2))
}
