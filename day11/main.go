package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

func main() {
	// Part 1: ./day11 -ngen 25 : 199946
    // Part 2: ./day11 -ngen 75 : 237994815702032

	ngen := flag.Int("ngen", 25, "number of generations to calculate")
	pdata := flag.String("data", "", "Stones like '125 17 3'")
	flag.Parse()

	var data string
	if *pdata == "" {
		bytes, err := os.ReadFile("testdata.dat")
		if err != nil {
			panic(err)
		}
		raw := string(bytes)
		data = string(raw)
	} else {
		data = *pdata
	}
	firstGen := strings.Split(data, " ")

	start := time.Now()

	n := dedup(&firstGen).DoNgenMap(*ngen)
	duration := time.Since(start)
	slog.Info(fmt.Sprintf("Stones: %d  ---  run duration: %v", n, duration))
}
