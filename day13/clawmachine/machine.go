package clawmachine

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// --------------------------------------------
//   Machine
// --------------------------------------------

type Machine struct {
	btnA        Button
	btnB        Button
	armPosition Location
	prizeAt     Location
}

func (machine Machine) String() string {
	return fmt.Sprintf("Prize at %s\n  %s\n  %s", machine.prizeAt, machine.btnA, machine.btnB)
}

func (m *Machine) IncreasePrizeLocationForPart2(n int) {
    m.prizeAt.x += n
    m.prizeAt.y += n
}

func LowestCost(totalPressesA []int, totalPressesB []int) (int, *string) {
	if len(totalPressesA) == 0 {
		err := "prize not grabbed"
		return 0, &err
	}

	minCost := math.MaxInt
	for i, pressA := range totalPressesA {
		cost := cost(pressA, totalPressesB[i])
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost, nil
}

func cost(pressesA int, pressesB int) int {
	return pressesA*3 + pressesB*1
}

func (m *Machine) FindPrize() ([]int, []int) {
	// Do 1 step button A then
	//   do many steps Button B until x or y is > than prize
	totalPressesA := []int{}
	totalPressesB := []int{}
	pressesA := 0
	pressesB := 0
LoopA:
	for atimes := 1; !m.armPosition.IsPast(m.prizeAt); atimes++ {
		m.pressA(atimes)
		if m.armPosition.IsPast(m.prizeAt) {
			break LoopA
		}
		pressesA++
		pressesB = 0
	LoopB:
		for pressB := 0; !m.armPosition.IsPast(m.prizeAt); pressB++ {
			pressesB++
			m.pressB(1)
			if m.IsArmOverPrize() {
				totalPressesA = append(totalPressesA, pressesA)
				totalPressesB = append(totalPressesB, pressesB)
				break LoopB
			}
		}
		m.resetArm()
	}
	return totalPressesA, totalPressesB
}

func (m *Machine) moveArm(to Location) {
	m.armPosition = to
}

func (m *Machine) resetArm() {
	m.armPosition = Location{}
	m.btnA.pressCount = 0
	m.btnB.pressCount = 0
}

func (m Machine) IsArmOverPrize() bool {
	return m.armPosition == m.prizeAt
}

func (m *Machine) pressA(times int) {
	m.btnA.pressCount += times
	m.moveArm(Location{
		x: m.armPosition.x + times*m.btnA.xOffset,
		y: m.armPosition.y + times*m.btnA.yOffset,
	})
}

func (m *Machine) pressB(times int) {
	m.btnB.pressCount += times
	m.moveArm(Location{
		x: m.armPosition.x + times*m.btnB.xOffset,
		y: m.armPosition.y + times*m.btnB.yOffset,
	})
}

func MachinefromStr(s string) Machine {
	machine := new(Machine)
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		if line[:6] == "Button" {
			data := line[10:]
			parts := strings.Split(data, ",")
			xData := strings.Split(strings.Trim(parts[0], " "), "+")[1]
			yData := strings.Split(strings.Trim(parts[1], " "), "+")[1]
			x, _ := strconv.Atoi(xData)
			y, _ := strconv.Atoi(yData)
			name := line[7:8]
			if name == "A" {
				machine.btnA = Button{name, 0, x, y}
			} else {
				machine.btnB = Button{name, 0, x, y}
			}
		}
		if line[:6] == "Prize:" {
			data := line[7:]
			parts := strings.Split(data, ",")
			xData := strings.Split(strings.Trim(parts[0], " "), "=")[1]
			yData := strings.Split(strings.Trim(parts[1], " "), "=")[1]
			x, _ := strconv.Atoi(xData)
			y, _ := strconv.Atoi(yData)
			machine.prizeAt = Location{x, y}
		}
	}
	return *machine
}

// --------------------------------------------
//   Button
// --------------------------------------------

type Button struct {
	name       string
	pressCount int
	xOffset    int
	yOffset    int
}

func (btn Button) String() string {
	return fmt.Sprintf("Button %s. Offsets x=%d, y=%d", btn.name, btn.xOffset, btn.yOffset)
}

func (btn *Button) Press() {
	btn.pressCount++
}

func (btn Button) PressCount() int {
	return btn.pressCount
}

func (btn Button) XOffset() int {
	return btn.xOffset
}

func (btn Button) YOffset() int {
	return btn.yOffset
}

func (btn Button) Position() Location {
	return Location{btn.xOffset * btn.pressCount, btn.yOffset * btn.pressCount}
}
