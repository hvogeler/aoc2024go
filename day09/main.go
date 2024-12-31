package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("example1.dat")
	data := string(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data: ", data)
}

type DiskMap string

type Disk struct {
	freeBlocks []Free
	files      []File
}

func (disk Disk) String() string {
	return fmt.Sprintf("Disk Layout:\n   Free Blocks: %v\n   File Blocks: %v", disk.freeBlocks, disk.files)
}

func FromDiskMap(diskMap *DiskMap) Disk {
	disk := Disk{}
	blocksizes := toNumberSlice(diskMap)
	currentPosition := uint(0)
	id := uint(0)
	for i, blocksize := range blocksizes {
		if i%2 == 1 {
			// free blocks
			disk.freeBlocks = append(disk.freeBlocks, Free{currentPosition, blocksize})
		} else {
			// file blocks
			disk.files = append(disk.files, File{id, currentPosition, blocksize})
			id++
		}
		currentPosition += blocksize
	}
	return disk
}

func toNumberSlice(diskMap *DiskMap) []uint {
	runes := []rune(*diskMap)
	var result = make([]uint, 0)
	for mapIdx := 0; mapIdx < len(runes); mapIdx++ {
		n := uint(runes[mapIdx] - 48)
		result = append(result, n)
	}
	return result
}

type File struct {
	id     uint
	start  uint // start of the block
	blocks uint // disk blocks occupied by file
}

type Free struct {
	start  uint
	blocks uint
}
