package warehouse2

import (
	"bufio"
	wh1 "day15/warehouse"
	"strings"
)

type Warehouse struct {
	robot Robot
	items map[wh1.Location]Item
	robotPath wh1.Path
	dimensions Dimensions
}

func NewWarehouse() Warehouse {
	wh := new(Warehouse)
	wh.items = make(map[wh1.Location]Item)
	return *wh
}

func WarehouseFromStr(s string) Warehouse {
	wh := NewWarehouse()
	scanner := bufio.NewScanner(strings.NewReader(s))
	y := 0
WarehouseLoop:
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break WarehouseLoop
		}
		if wh.dimensions.x == 0 {
			wh.dimensions.x = len(line)*2
		}
		runes := []rune(line)
		xOut := 0
		for x := 0; x < len(line); x++ {
			rne := runes[x]
			switch rne {
			case WallRune:
				wall := NewWall(x, y)
				wh.items[wh1.NewLocation(x, y)] = wall
				wh.items[wh1.NewLocation(x + 1, y)] = wall
				xOut += 2
			case BoxRune:
				box := NewBox(x, y)
				wh.items[wh1.NewLocation(x, y)] = box
				wh.items[wh1.NewLocation(x + 1, y)] = box
				xOut += 2
			case RobotRune:
				wh.robot = NewRobot(x, y)
				xOut += 2
			case UnusedRune:
				xOut += 2
			default:
				panic("Exhausted Switch")
			}
		}
		y++
	}
	wh.dimensions.x *= 2
	wh.dimensions.y = y

	for scanner.Scan() {
		line := scanner.Text()
		for _, pointer := range line {
			wh.robotPath.AddPointer(wh1.Pointer(pointer))
		}
	}
	return wh
}


type Dimensions struct {
	x int
	y int
}