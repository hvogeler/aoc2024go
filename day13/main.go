package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("example.dat")
	if err != nil {
		panic(err)
	}
	data := string(bytes)
	fmt.Println(data)
}
