package main

import (
	"fmt"
	"testing"
)

const good1 = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`

const good2 = `Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450`

const bad1 = `Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176`

const bad2 = `Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=1027`

func Test_t1(t *testing.T) {
	t.Run("read good1", func(t *testing.T) {
		fmt.Printf("TEST")
	})

}
