package clawmachine

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Machine struct {
	btnA    Button
	btnB    Button
	prizeAt Location
}

func (machine Machine) String() string {
	return fmt.Sprintf("Prize at %s\n  %s\n  %s", machine.prizeAt, machine.btnA, machine.btnB)
}

type Button struct {
	name    string
	xOffset int
	yOffset int
}

func (btn Button) String() string {
	return fmt.Sprintf("Button %s. Offsets x=%d, y=%d", btn.name, btn.xOffset, btn.yOffset)
}

func machinefromStr(s string) Machine {
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
				machine.btnA = Button{name, x, y}
			} else {
				machine.btnB = Button{name, x, y}
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
