package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("testdata.dat")
	if err != nil {
		panic(err)
	}
	data := string(bytes)
	fmt.Println(data)
}
