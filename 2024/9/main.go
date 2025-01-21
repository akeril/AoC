package main

import (
	"cmp"
	"fmt"

	"github.com/akeril/aoc2024/utils"
)

func main() {
	input := utils.ReadFile("input")[0]
	fs := ParseInput(input)
	fs = CompactFS(fs)
	fmt.Println(fs.CheckSum())
}

type Block struct {
	pos, id, freq int
}

func (b Block) Compare(o Block) int {
	return cmp.Compare(b.pos, o.pos)
}

type FileSystem struct {
	FileBlocks []Block
	FreeBlocks []Block
}

func (fs *FileSystem) CheckSum() int {
	checksum := 0
	for _, block := range fs.FileBlocks {
		for k := 0; k < block.freq; k++ {
			checksum += block.id * (block.pos + k)
		}
	}
	return checksum
}

func ParseInput(input string) FileSystem {
	p := 0
	i := 0
	free := false
	fs := FileSystem{FileBlocks: make([]Block, 0), FreeBlocks: make([]Block, 0)}
	for _, c := range input {
		k := int(c - '0')
		if free {
			fs.FreeBlocks = append(fs.FreeBlocks, Block{pos: p, freq: k})
			p += k
		} else {
			fs.FileBlocks = append(fs.FileBlocks, Block{pos: p, id: i, freq: k})
			p += k
			i++
		}
		free = !free
	}
	return fs
}

func CompactFS(fs FileSystem) FileSystem {
	for j := len(fs.FileBlocks) - 1; j >= 0; j-- {
		// find valid block
		for i := 0; i < len(fs.FreeBlocks); i++ {
			if fs.FreeBlocks[i].pos < fs.FileBlocks[j].pos && fs.FreeBlocks[i].freq >= fs.FileBlocks[j].freq {
				fs.FileBlocks[j].pos = fs.FreeBlocks[i].pos
				fs.FreeBlocks[i].pos += fs.FileBlocks[j].freq
				fs.FreeBlocks[i].freq -= fs.FileBlocks[j].freq
				break
			}
		}
	}
	return fs
}
