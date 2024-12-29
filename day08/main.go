package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("example.dat")
    rawMap := string(bytes)
    if err != nil {
        panic(err)
    }

    cityMap := cityMapFromStr(rawMap)
    fmt.Println(cityMap)
}
