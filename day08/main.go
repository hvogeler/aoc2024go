package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("example.dat")
    data1 := string(data)
    if err != nil {
        panic(err)
    }

	fmt.Println(data1)	
    scanner := bufio.NewScanner(strings.NewReader(data1))
    i := 0
    for scanner.Scan() {
        i++
        line := scanner.Text()
        fmt.Printf("%3d %s\n", i, line)
    }
}