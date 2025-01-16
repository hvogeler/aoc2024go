package clawmachine

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
		machine := machinefromStr(good1)
		fmt.Println(machine)
		if machine.prizeAt != (Location{8400, 5400}) {
			t.Errorf("Wrong location")
		}
		if machine.btnA.xOffset != 94 {
			t.Errorf("Wrong x offset button A")
		}
		if machine.btnB.xOffset != 22 {
			t.Errorf("Wrong x offset button B")
		}
		if machine.btnA.yOffset != 34 {
			t.Errorf("Wrong y offset button A: %d", machine.btnA.yOffset)
		}
		if machine.btnB.yOffset != 67 {
			t.Errorf("Wrong y offset button B: %d", machine.btnB.yOffset)
		}
	})

	t.Run("read bad1", func(t *testing.T) {
		machine := machinefromStr(bad1)
		fmt.Println(machine)
		if machine.prizeAt != (Location{12748, 12176}) {
			t.Errorf("Wrong location")
		}
		if machine.btnA.xOffset != 26 {
			t.Errorf("Wrong x offset button A")
		}
		if machine.btnB.xOffset != 67 {
			t.Errorf("Wrong x offset button B")
		}
		if machine.btnA.yOffset != 66 {
			t.Errorf("Wrong y offset button A: %d", machine.btnA.yOffset)
		}
		if machine.btnB.yOffset != 21 {
			t.Errorf("Wrong y offset button B: %d", machine.btnB.yOffset)
		}
	})
}
