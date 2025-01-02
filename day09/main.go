package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("testdata.dat")
	data := DiskMap(bytes)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Data: ", data)

	disk := FromDiskMap(&data)
	disk.Compress1()
	cs := disk.Checksum()
	fmt.Printf("Part1 Checksum: %d\n", cs)
	if cs != 6360094256423 {
		panic("Wrong result")
	}

	// Part 2
	disk = FromDiskMap(&data)
	disk.Compress2()
	cs = disk.Checksum()
	fmt.Printf("Part2 Checksum: %d\n", cs)
	if cs != 6379677752410 {
		panic("Wrong result")
	}

}
