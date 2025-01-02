package main

import (
	"fmt"
	"testing"
)

func Test_Disk(t *testing.T) {
	t.Run("FromDiskMap", func(t *testing.T) {
		testdata := DiskMap("2333133121414131402")
		disk := FromDiskMap(&testdata)
		fmt.Println("Before : ", disk)
		disk.Compress()
		fmt.Println("After  : ", disk)
		cs := disk.Checksum()
		fmt.Println("Checksum = ", cs)
	})

	t.Run("FromDiskMap", func(t *testing.T) {
		testdata := DiskMap("12345")
		disk := FromDiskMap(&testdata)
		if len(disk.fileBlocks) != 3 {
			t.Errorf("Expected 3 files, got %d", len(disk.fileBlocks))
		}
		if len(disk.freeBlocks) != 2 {
			t.Errorf("Expected 2 free blocks, got %d", len(disk.freeBlocks))
		}
		f := BlockRange{File, 2, 10, 5}
		if disk.fileBlocks[2] != f {
			t.Errorf("Expected file %v, got %v", f, disk.fileBlocks[2])
		}
		if disk.SpaceFree() != 6 {
			t.Errorf("Expected 6 blocks free space, got %d", disk.SpaceFree())
		}
		if disk.SpaceUsed() != 9 {
			t.Errorf("Expected 9 blocks used space, got %d", disk.SpaceUsed())
		}
		if disk.SpaceTotal() != 15 {
			t.Errorf("Expected 15 blocks total disk space, got %d", disk.SpaceTotal())
		}
		fmt.Println(disk)
		disk.Compress()
		fmt.Println(disk)

	})

	t.Run("runeToNumber", func(t *testing.T) {
		testdata := DiskMap("12345")
		a := toNumberSlice(&testdata)
		var array [5]uint
		copy(array[:], a)
		if array != [5]uint{1, 2, 3, 4, 5} {
			t.Errorf("toNumberSlice failed")
		}

	})
}
