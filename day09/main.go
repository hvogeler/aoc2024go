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
	disk.Compress()
	cs := disk.Checksum()
	fmt.Printf("Part1 Checksum: %d", cs)
	if cs != 6360094256423 {
		panic("Wrong result")
	}

}
