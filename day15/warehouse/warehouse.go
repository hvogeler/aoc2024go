package warehouse

import (
	"bufio"
	"strings"
)

type Warehouse struct {
	robot      Robot
	boxes      map[Location]Box
	walls      map[Location]Wall
	robotPath  Path
	dimensions Dimensions
}

func NewWarehouse() Warehouse {
	wh := new(Warehouse)
	wh.boxes = make(map[Location]Box)
	wh.walls = make(map[Location]Wall)
	return *wh
}

func (wh Warehouse) String() string {
	s := ""
	for y := 0; y < wh.dimensions.y; y++ {
		for x := 0; x < wh.dimensions.x; x++ {
			s += wh.ObjectTypeAt(NewLocation(x, y)).String()
		}
		s += "\n"
	}
	return s
}

func (wh Warehouse) ObjectTypeAt(position Location) ObjectType {
	if _, exists := wh.boxes[position]; exists {
		return BoxType
	}
	if _, exists := wh.walls[position]; exists {
		return WallType
	}
	if wh.robot.position == position {
		return RobotType
	}
	return UnusedType
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
			wh.dimensions.x = len(line)
		}
		runes := []rune(line)
		for x, rne := range runes {
			switch rne {
			case WallRune:
				wh.walls[NewLocation(x, y)] = NewWall(x, y)
			case BoxRune:
				wh.boxes[NewLocation(x, y)] = NewBox(x, y)
			case RobotRune:
				wh.robot = NewRobot(x, y)
			}
		}
		y++
	}
	wh.dimensions.y = y

	for scanner.Scan() {
		line := scanner.Text()
		pointers := []rune(line)
		for _, pointer := range pointers {
			wh.robotPath.pointers = append(wh.robotPath.pointers, Pointer(pointer))
		}
	}
	return wh
}

type Dimensions struct {
	x int
	y int
}

const UnusedRune rune = '.'

type ObjectType rune

const (
	WallType   ObjectType = ObjectType(WallRune)
	BoxType    ObjectType = ObjectType(BoxRune)
	RobotType  ObjectType = ObjectType(RobotRune)
	UnusedType ObjectType = ObjectType(UnusedRune)
)

func (ot ObjectType) String() string {
	return string(ot)
}
