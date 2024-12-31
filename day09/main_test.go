package main

import (
	"testing"
	"fmt"
)

func Test_Disk(t *testing.T) {

	t.Run("FromDiskMap", func(t *testing.T) {
		testdata := DiskMap("12345")
		disk := FromDiskMap(&testdata)
		if len(disk.files) != 3 {
			t.Errorf("Expected 3 files, got %d", len(disk.files))
		}
		if len(disk.freeBlocks) != 2 {
			t.Errorf("Expected 2 free blocks, got %d", len(disk.freeBlocks))
		}
		f := File{2, 10, 5}
		if disk.files[2] != f {
			t.Errorf("Expected file %v, got %v", f, disk.files[2])
		}
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
