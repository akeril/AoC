package main

import (
	"cmp"
	"fmt"

	"github.com/kjabin/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")[0]
	fs := ParseInput(input)
	fs = CompactFS(fs)
	fmt.Println(fs.CheckSum())
}

type Block struct {
	pos, id int
}

func (b Block) Compare(o Block) int {
	return cmp.Compare(b.pos, o.pos)
}

type FileSystem struct {
	FileBlocks []Block
	FreeBlocks []int
}

func (fs *FileSystem) CheckSum() int {
	checksum := 0
	for _, block := range fs.FileBlocks {
		checksum += block.id * block.pos
	}
	return checksum
}

func ParseInput(input string) FileSystem {
	p := 0
	i := 0
	free := false
	fs := FileSystem{FileBlocks: make([]Block, 0), FreeBlocks: make([]int, 0)}
	for _, c := range input {
		k := int(c - '0')
		if free {
			for k > 0 {
				fs.FreeBlocks = append(fs.FreeBlocks, p)
				p++
				k--
			}
		} else {
			for k > 0 {
				fs.FileBlocks = append(fs.FileBlocks, Block{pos: p, id: i})
				p++
				k--
			}
			i++
		}
		free = !free
	}
	return fs
}

func CompactFS(fs FileSystem) FileSystem {
	i := 0
	j := len(fs.FileBlocks) - 1

	for fs.FreeBlocks[i] < fs.FileBlocks[j].pos {
		fs.FileBlocks[j].pos, fs.FreeBlocks[i] = fs.FreeBlocks[i], fs.FileBlocks[j].pos
		i++
		j--
	}

	return fs
}
