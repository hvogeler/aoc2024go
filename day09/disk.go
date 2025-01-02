package main

import (
	"fmt"
	"sort"
)

type DiskMap string

type Disk struct {
	freeBlocks []BlockRange
	fileBlocks []BlockRange
}

func (disk *Disk) Checksum() uint {
	checkSum := uint(0)
	for _, file := range disk.fileBlocks {
		for blockPos := file.start; blockPos < file.start+file.blocks; blockPos++ {
			checkSum += (uint(blockPos) * file.id)
		}
	}
	return checkSum
}

// move file blocks one at a time from the end of the disk to the leftmost free space block
// until there are no gaps remaining between file blocks
func (disk *Disk) Compress() {
	for disk.isFreeBlockBetweenFiles() {
		lastFileIdx := len(disk.fileBlocks) - 1
		lastFileBlocks := disk.fileBlocks[lastFileIdx].blocks
		firstFreeBlocks := disk.freeBlocks[0].blocks
		switch {
		case lastFileBlocks >= firstFreeBlocks:
			disk.moveLastFileToFreePosition()
			disk.removeEmptyLastFileBlock()
			disk.moveFreeBlocksToEnd()
		case lastFileBlocks < firstFreeBlocks:
			disk.freeBlocks[0].blocks -= disk.fileBlocks[lastFileIdx].blocks
			disk.freeBlocks[len(disk.freeBlocks)-1].blocks += disk.fileBlocks[lastFileIdx].blocks
			disk.fileBlocks[lastFileIdx].start = disk.freeBlocks[0].start
			disk.freeBlocks[0].start += (disk.fileBlocks[lastFileIdx].blocks)
			disk.removeEmptyLastFileBlock()
			disk.SortFilesByStart()
		}
	}
}

func (disk *Disk) moveFreeBlocksToEnd() {
	disk.freeBlocks[len(disk.freeBlocks)-1].blocks += disk.freeBlocks[0].blocks
	disk.freeBlocks = disk.freeBlocks[1:]
}

func (disk *Disk) moveLastFileToFreePosition() {
	to := 0
Loop:
	for ; to < len(disk.fileBlocks); to++ {
		if disk.freeBlocks[0].start < disk.fileBlocks[to].start {
			break Loop
		}
	}

	lastFileIdx := len(disk.fileBlocks) - 1
	disk.fileBlocks[lastFileIdx].blocks -= disk.freeBlocks[0].blocks
	disk.fileBlocks = append(disk.fileBlocks[:to], append([]BlockRange{{File, disk.fileBlocks[lastFileIdx].id, disk.freeBlocks[0].start, disk.freeBlocks[0].blocks}}, disk.fileBlocks[to:]...)...)
}

// func (disk *Disk) moveLastFileTo(to int) {
// 	lastFileIdx := len(disk.fileBlocks) - 1
// 	disk.fileBlocks[lastFileIdx].blocks -= disk.freeBlocks[0].blocks
// 	disk.fileBlocks = append(disk.fileBlocks[:to], append([]File{{disk.fileBlocks[lastFileIdx].id, disk.freeBlocks[0].start, disk.freeBlocks[0].blocks}}, disk.fileBlocks[to:]...)...)
// }

func (disk *Disk) removeEmptyLastFileBlock() {
	lastFileIdx := len(disk.fileBlocks) - 1
	if disk.fileBlocks[lastFileIdx].blocks == 0 {
		disk.fileBlocks = disk.fileBlocks[:len(disk.fileBlocks)-1]
	}
}

func (disk Disk) isFreeBlockBetweenFiles() bool {
	return disk.minStartFreeBlock() < disk.maxStartFileBlock()
}

// func (disk Disk) isFileBlock(blockPos uint) bool {
// 	for _, fileBlock := range disk.fileBlocks {
// 		if blockPos >= fileBlock.start && blockPos < fileBlock.start+fileBlock.blocks {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (disk Disk) maxStartFreeBlock() uint {
// 	max := uint(0)
// 	for _, freeBlock := range disk.freeBlocks {
// 		if freeBlock.start > max {
// 			max = freeBlock.start
// 		}
// 	}
// 	return max
// }

func (disk Disk) minStartFreeBlock() uint {
	min := ^uint(0)
	for _, freeBlock := range disk.freeBlocks {
		if freeBlock.start < min {
			min = freeBlock.start
		}
	}
	return min
}

func (disk Disk) maxStartFileBlock() uint {
	max := uint(0)
	for _, fileBlock := range disk.fileBlocks {
		if fileBlock.start > max {
			max = fileBlock.start
		}
	}
	return max
}

func (disk *Disk) SortFilesByStart() {
	sort.Slice(disk.fileBlocks, func(i, j int) bool {
		return disk.fileBlocks[i].start < disk.fileBlocks[j].start
	})
}

func (disk Disk) SpaceUsed() uint {
	usedSpace := uint(0)
	for _, file := range disk.fileBlocks {
		usedSpace += file.blocks
	}
	return usedSpace
}

func (disk Disk) SpaceFree() uint {
	freeSpace := uint(0)
	for _, free := range disk.freeBlocks {
		freeSpace += free.blocks
	}
	return freeSpace
}

func (disk Disk) SpaceTotal() uint {
	return disk.SpaceFree() + disk.SpaceUsed()
}

func (disk Disk) String() string {
	allBlocks := append(disk.fileBlocks, disk.freeBlocks...)
	sort.Slice(allBlocks, func(i, j int) bool {
		return allBlocks[i].start < allBlocks[j].start
	})
	diskMap := ""
	for _, block := range allBlocks {
		for i := 0; i < int(block.blocks); i++ {
			if block.usage == File {
				diskMap += fmt.Sprintf("%d", block.id)
			} else {
				diskMap += "."
			}
		}
	}
	return diskMap
	// return fmt.Sprintf("Disk Layout:\n   Free Blocks: %v\n   File Blocks: %v", disk.freeBlocks, disk.fileBlocks)
}

func FromDiskMap(diskMap *DiskMap) Disk {
	disk := Disk{}
	blocksizes := toNumberSlice(diskMap)
	currentPosition := uint(0)
	id := uint(0)
	for i, blocksize := range blocksizes {
		if i%2 == 1 {
			// free blocks
			disk.freeBlocks = append(disk.freeBlocks, BlockRange{Empty, EmptyBlockId, currentPosition, blocksize})
		} else {
			// file blocks
			disk.fileBlocks = append(disk.fileBlocks, BlockRange{File, id, currentPosition, blocksize})
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

type Usage uint

const (
	Empty Usage = iota
	File
)

func (usage Usage) String() string {
	switch usage {
	case Empty:
		return "Empty"
	case File:
		return "File"
	}
	panic("Unknown Block usage")
}

const EmptyBlockId = uint(0)

type BlockRange struct {
	usage  Usage
	id     uint
	start  uint // start of the block
	blocks uint // disk blocks occupied by file
}

func (f BlockRange) String() string {
	return fmt.Sprintf("Type(%s) ID(%2d) Start(%3d) Blocks(%3d)\n", f.usage, f.id, f.start, f.blocks)
}
