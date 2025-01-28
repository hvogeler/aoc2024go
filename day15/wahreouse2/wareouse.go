package warehouse2

import (
	"bufio"
	wh1 "day15/warehouse"
	"fmt"
	"strings"
)

type Warehouse struct {
	robot       Robot
	items       map[Location]Item
	robotPath   wh1.Path
	dimensions  Dimensions
	undoLog     []UndoItem
	unwindCause ItemType
}

func NewWarehouse() Warehouse {
	wh := new(Warehouse)
	wh.items = make(map[Location]Item)
	return *wh
}

func (wh *Warehouse) Move(itemAt *Location, direction wh1.Pointer, level int) {
	if level == 0 {
		wh.ClearUndoLog()
	}
	level++
	currentItem, curExists := wh.ItemAtPosition(*itemAt)
	if !curExists {
		panic("Current Object for Move must exist")
	}
	nextPositionMove := new(Location)
	nextPositionTest := new(Location)
	switch direction {
	case wh1.Right:
		nextPositionMove.SetX(currentItem.PositionLeft().X() + 1)
		nextPositionMove.SetY(currentItem.PositionLeft().Y())
		nextPositionTest.SetX(currentItem.PositionLeft().X() + currentItem.Length())
		nextPositionTest.SetY(currentItem.PositionLeft().Y())
	case wh1.Left:
		nextPositionMove.SetX(currentItem.PositionLeft().X() - 1)
		nextPositionMove.SetY(currentItem.PositionLeft().Y())
		nextPositionTest.SetX(currentItem.PositionLeft().X() - 1) //currentItem.Length())
		nextPositionTest.SetY(currentItem.PositionLeft().Y())
	case wh1.Up:
		nextPositionMove.SetX(currentItem.PositionLeft().X())
		nextPositionMove.SetY(currentItem.PositionLeft().Y() - 1)
		nextPositionTest.SetX(currentItem.PositionLeft().X())
		nextPositionTest.SetY(currentItem.PositionLeft().Y() - 1)
	case wh1.Down:
		nextPositionMove.SetX(currentItem.PositionLeft().X())
		nextPositionMove.SetY(currentItem.PositionLeft().Y() + 1)
		nextPositionTest.SetX(currentItem.PositionLeft().X())
		nextPositionTest.SetY(currentItem.PositionLeft().Y() + 1)
	default:
		panic("Exhausted switch")
	}

	// Process Wall ahead
	nextItem, exists := wh.ItemAtPosition(*nextPositionTest)
	nextItemRight, existsRight := wh.ItemAtPosition(nextPositionTest.Right())
	if exists && existsRight {
		if nextItemRight.PositionLeft() == nextItem.PositionLeft() {
			existsRight = false
		}
	}
	switch direction {
	case wh1.Left, wh1.Right:
		if exists && nextItem.Item() == WallItem {
			wh.unwindCause = WallItem
			return
		}
	case wh1.Up, wh1.Down:
		if exists && nextItem.Item() == WallItem {
			wh.unwindCause = WallItem
			return
		}
		if existsRight && nextItemRight.Item() == WallItem {
			wh.unwindCause = WallItem
			return
		}
	}

	// Process Box ahead
	if exists && nextItem.Item() == BoxItem && direction.Orientation() == wh1.Horizontal {
		wh.Move(nextPositionTest, direction, level)
	}

	if (exists && nextItem.Item() == BoxItem || existsRight && nextItemRight.Item() == BoxItem) && direction.Orientation() == wh1.Vertical {
		switch currentItem.Item() {
		case RobotItem:
			if exists {
				nextLeft := nextItem.PositionLeft()
				wh.Move(&nextLeft, direction, level)
			}
		case BoxItem:
			if exists {
				partLeft := nextItem.PositionLeft()
				wh.Move(&partLeft, direction, level)
			}
			if existsRight {
				partLeft := nextItemRight.PositionLeft()
				wh.Move(&partLeft, direction, level)
			}
		}
	}

	if level == 1 && wh.unwindCause == WallItem {
		fmt.Println("**** Hit a Wall ****", level)
		fmt.Println(wh.undoLog)
		wh.Undo()
		return
	}

	if wh.unwindCause != WallItem {
		switch currentItem.Item() {
		case BoxItem:
			locationToDelete := currentItem.PositionLeft()
			currentItem.SetPosition(*nextPositionMove)
			wh.items[*nextPositionMove] = currentItem
			wh.AddUndoItem(&currentItem, direction)
			delete(wh.items, locationToDelete)
			return
		case RobotItem:
			locationToDelete := wh.robot.PositionLeft()
			wh.robot.position = *nextPositionMove
			wh.items[*nextPositionMove] = &wh.robot
			delete(wh.items, locationToDelete)
			return
		}
	}
}

func (wh *Warehouse) GoRobotGo() {
	for i := 0; i < wh.robotPath.Length(); i++ {
		ptr := wh.robotPath.NextPointer()
		wh.Move(&wh.robot.position, ptr, 0)
	}
}

func (wh Warehouse) SumBoxCoords() int {
	sum := 0
	for _, v := range wh.items {
		if v.Item() == BoxItem {
			gpsCoord := v.PositionLeft().y*100 + v.PositionLeft().x
			sum += gpsCoord
		}
	}
	return sum
}

func (wh *Warehouse) AddUndoItem(item *Item, pointer wh1.Pointer) {
	wh.undoLog = append(wh.undoLog, UndoItem{item, pointer})
}

func (wh *Warehouse) ClearUndoLog() {
	wh.undoLog = wh.undoLog[:0]
	wh.unwindCause = UnusedItem
}

func (wh *Warehouse) Undo() {
	for _, undoItem := range wh.undoLog {
		step := 1
		if undoItem.direction == wh1.Down || undoItem.direction == wh1.Right {
			step = -1
		}
		locationToDelete := (*undoItem.item).PositionLeft()
		originalPosition := NewLocation(locationToDelete.x, locationToDelete.y+step)
		if undoItem.direction == wh1.Left || undoItem.direction == wh1.Right {
			originalPosition = NewLocation(locationToDelete.x+step, locationToDelete.y)
		}
		(*undoItem.item).SetPosition(originalPosition)
		wh.items[originalPosition] = *undoItem.item
		delete(wh.items, locationToDelete)
		wh.ClearUndoLog()
	}
}

func (wh Warehouse) ItemAtPosition(loc Location) (Item, bool) {
	return wh.ItemAt(loc.X(), loc.Y())
}

func (wh Warehouse) ItemAt(x, y int) (Item, bool) {
	item, exists := wh.items[NewLocation(x, y)]
	if exists {
		return item, true
	}
	item, exists = wh.items[NewLocation(x-1, y)]
	if exists && item.Length() == 2 {
		return item, true
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
				wh.items[NewLocation(xOut, y)] = &wall
			case BoxRune:
				box := NewBox(xOut, y)
				wh.items[NewLocation(xOut, y)] = &box
			case RobotRune:
				wh.robot = NewRobot(xOut, y)
				wh.items[NewLocation(xOut, y)] = &wh.robot
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
				item := itemPtr
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
