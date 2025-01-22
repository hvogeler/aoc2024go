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

func (wh *Warehouse) Move(objectAt *Location, direction Pointer) {
	nextPosition := new(Location)
	switch direction {
	case right:
		nextPosition.x = objectAt.x + 1
		nextPosition.y = objectAt.y
	case left:
		nextPosition.x = objectAt.x - 1
		nextPosition.y = objectAt.y
	case up:
		nextPosition.x = objectAt.x
		nextPosition.y = objectAt.y - 1
	case down:
		nextPosition.x = objectAt.x
		nextPosition.y = objectAt.y + 1
	default:
		panic("Exhausted switch")
	}

	if wh.ObjectTypeAt(*nextPosition) == WallType {
		return
	}

	if wh.ObjectTypeAt(*nextPosition) == BoxType {
		wh.Move(nextPosition, direction)
	}

	if wh.ObjectTypeAt(*nextPosition) == UnusedType {
		switch wh.ObjectTypeAt(*objectAt) {
		case BoxType:
			box := wh.boxes[*objectAt]
			box.position = *nextPosition
			wh.boxes[*nextPosition] = box
			delete(wh.boxes, *objectAt)
			return
		case RobotType:
			wh.robot.position = *nextPosition
			return
		}
	}
}

func (wh *Warehouse) GoRobotGo() {
	for i := 0; i < wh.robotPath.Length(); i++ {
		ptr := wh.robotPath.NextPointer()
		wh.Move(&wh.robot.position, ptr)
	}
}

func (wh Warehouse) SumBoxCoords() int {
	sum := 0
	for _, v := range wh.boxes {
		sum += v.GpsCoord()
	}
	return sum
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
			switch string(rne) {
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
		for _, pointer := range line {
			wh.robotPath.pointers = append(wh.robotPath.pointers, Pointer(pointer))
		}
	}
	return wh
}

type Dimensions struct {
	x int
	y int
}

const UnusedRune string = "."

type ObjectType string

const (
	WallType   ObjectType = ObjectType(WallRune)
	BoxType    ObjectType = ObjectType(BoxRune)
	RobotType  ObjectType = ObjectType(RobotRune)
	UnusedType ObjectType = ObjectType(UnusedRune)
)

func (ot ObjectType) String() string {
	return string(ot)
}
