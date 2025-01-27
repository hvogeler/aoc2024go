package warehouse2

import (
	"bufio"
	wh1 "day15/warehouse"
	"strings"
)

type Warehouse struct {
	robot      Robot
	items      map[wh1.Location]Item
	robotPath  wh1.Path
	dimensions Dimensions
}

func NewWarehouse() Warehouse {
	wh := new(Warehouse)
	wh.items = make(map[wh1.Location]Item)
	return *wh
}

func (wh Warehouse) ItemAt(x, y int) (*Item, bool) {
	item, exists := wh.items[wh1.NewLocation(x, y)]
	if exists {
		return &item, true
	}
	item, exists = wh.items[wh1.NewLocation(x - 1, y)]
    if exists && item.Length() == 2{
		return &item, true
	}
	return nil, false
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
			wh.dimensions.x = len(line) * 2
		}
		runes := []rune(line)
		xOut := 0
		for x := 0; x < len(line); x++ {
			rne := runes[x]
			switch rne {
			case WallRune:
				wall := NewWall(xOut, y)
				wh.items[wh1.NewLocation(xOut, y)] = wall
			case BoxRune:
				box := NewBox(xOut, y)
				wh.items[wh1.NewLocation(xOut, y)] = box
			case RobotRune:
				wh.robot = NewRobot(xOut, y)
				wh.items[wh1.NewLocation(xOut, y)] = wh.robot
			case UnusedRune:
			default:
				panic("Exhausted Switch")
			}
            xOut += 2
		}
		y++
	}
	wh.dimensions.y = y

	for scanner.Scan() {
		line := scanner.Text()
		for _, pointer := range line {
			wh.robotPath.AddPointer(wh1.Pointer(pointer))
		}
	}
	return wh
}

func (wh Warehouse) String() string {
	var builder strings.Builder
	for y := 0; y < wh.dimensions.y; y++ {
		for x := 0; x < wh.dimensions.x; {
			itemPtr, exists := wh.ItemAt(x, y)
			if exists {
				item := *itemPtr
				builder.WriteString(item.String())
                x += item.Length()
			} else {
				builder.WriteString(".")
                x++
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

type Dimensions struct {
	x int
	y int
}
