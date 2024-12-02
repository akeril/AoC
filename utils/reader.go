package utils

import (
	"bufio"
	"os"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(file)

	var text []string
	for sc.Scan() {
		text = append(text, sc.Text())
	}
	return text
}
