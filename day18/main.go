package main

import (
	ms "day18/memspace"
	"fmt"
)

func main() {
	data := ms.ReadData("example1.dat")
	fmt.Println(data)
}

