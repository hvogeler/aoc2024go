package main

import (
	"fmt"
	"day14/robots"
)

func main() {
	bytes := robots.ReadData("example.dat")
	data := string(bytes)
	fmt.Println(data)

}

