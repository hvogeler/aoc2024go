package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	bytes, err := os.ReadFile("testdata.dat")
	data := DiskMap(bytes)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Data: ", data)

	disk := FromDiskMap(&data)
	disk.Compress()
	cs := disk.Checksum()
	fmt.Printf("Part1 Checksum: %d", cs)
    if cs != 6360094256423 {
        panic("Wrong result")
    }

}

type DiskMap string

type Disk struct {
	freeBlocks []Free
	fileBlocks []File
}

func (disk *Disk) Checksum() uint {
	checkSum := uint(0)
	for _, file := range disk.fileBlocks {
		for blockPos := file.start; blockPos < file.start + file.blocks; blockPos++ {
			checkSum += (uint(blockPos) * file.id)
		}
	}
	return checkSum
}

func (disk *Disk) Compress() {
	for disk.minStartFreeBlock() < disk.maxStartFileBlock() {
		lastFile := len(disk.fileBlocks) - 1
		if disk.freeBlocks[0].blocks <= disk.fileBlocks[lastFile].blocks {
			disk.fileBlocks[lastFile].blocks -= disk.freeBlocks[0].blocks
			disk.fileBlocks = append(disk.fileBlocks[:1], append([]File{{disk.fileBlocks[lastFile].id, disk.freeBlocks[0].start, disk.freeBlocks[0].blocks}}, disk.fileBlocks[1:]...)...)
            if disk.fileBlocks[lastFile].blocks == 0 {
                disk.fileBlocks = disk.fileBlocks[:len(disk.fileBlocks) - 1]
            }
            disk.freeBlocks[len(disk.freeBlocks) - 1].blocks += disk.freeBlocks[0].blocks
			disk.freeBlocks = disk.freeBlocks[1:]
		} else {
			disk.freeBlocks[0].blocks -= disk.fileBlocks[lastFile].blocks
            disk.freeBlocks[len(disk.freeBlocks) - 1].blocks += disk.fileBlocks[lastFile].blocks
			disk.fileBlocks[lastFile].start = disk.freeBlocks[0].start
			disk.freeBlocks[0].start += (disk.fileBlocks[lastFile].blocks)
            if disk.fileBlocks[lastFile].blocks == 0 {
                disk.fileBlocks = disk.fileBlocks[:len(disk.fileBlocks) - 1]
            }
		}
        // here for debugging. Move it out of the loop for prod
        disk.SortFilesByStart()
	}
}

func (disk Disk) isFileBlock(blockPos uint) bool {
    for _, fileBlock := range disk.fileBlocks {
        if blockPos >= fileBlock.start && blockPos < fileBlock.start + fileBlock.blocks {
            return true
        }
    }
    return false
}

func (disk Disk) maxStartFreeBlock() uint {
	max := uint(0)
	for _, freeBlock := range disk.freeBlocks {
		if freeBlock.start > max {
			max = freeBlock.start
		}
	}
	return max
}

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
	return fmt.Sprintf("Disk Layout:\n   Free Blocks: %v\n   File Blocks: %v", disk.freeBlocks, disk.fileBlocks)
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
			disk.fileBlocks = append(disk.fileBlocks, File{id, currentPosition, blocksize})
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

func (f File) String() string {
	return fmt.Sprintf("ID(%2d) Start(%3d) Blocks(%3d)\n", f.id, f.start, f.blocks)
}

type Free struct {
	start  uint
	blocks uint
}

func (free Free) String() string {
	return fmt.Sprintf("Start(%3d) Blocks(%3d)\n", free.start, free.blocks)
}
