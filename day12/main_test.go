package main

import (
	"fmt"
	"testing"
)

const example1 = `AAAA
BBCD
BBCC
EEEC`

const example2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

const example3 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func Test_t1(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		fmt.Printf("TEST")
	})

}
