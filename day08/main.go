package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	bytes, err := os.ReadFile("testdata.dat")
	rawMap := string(bytes)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	cityMap := CityMapFromStr(rawMap)
	fmt.Println(cityMap)
	cityMap.CreateAntinodes()
	duration := time.Since(start)
	fmt.Println(cityMap)
	fmt.Printf("Number of Antinodes: %d\n", len(cityMap.antinodes()))
	fmt.Printf("Took %v\n", duration)

}
