package warehouse2

import (
	"bufio"
	wh1 "day15/warehouse"
	"fmt"
	"strings"
)


type Warehouse struct {
	robot      Robot
	items      map[Location]Item
	robotPath  wh1.Path
	dimensions Dimensions
	undoLog    []UndoItem
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
    level++;
	currentItem, curExists := wh.ItemAtPosition(*itemAt)
	if !curExists {
		panic("Current Object for Move must exist")
	}
	nextPosition := new(Location)
	switch direction {
	case wh1.Right:
		nextPosition.SetX(itemAt.X() + currentItem.Length())
		nextPosition.SetY(itemAt.Y())
	case wh1.Left:
		nextPosition.SetX(itemAt.X() - currentItem.Length())
		nextPosition.SetY(itemAt.Y())
	case wh1.Up:
		nextPosition.SetX(itemAt.X())
		nextPosition.SetY(itemAt.Y() - 1)
	case wh1.Down:
		nextPosition.SetX(itemAt.X())
		nextPosition.SetY(itemAt.Y() + 1)
	default:
		panic("Exhausted switch")
	}

    // Check for Wall
	nextItem, exists := wh.ItemAtPosition(*nextPosition)
    nextItemRight, existsRight := wh.ItemAtPosition(nextPosition.Right())
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

	if exists && nextItem.Item() == BoxItem {
		switch direction {
		case wh1.Left, wh1.Right:
			wh.Move(nextPosition, direction, level)
		case wh1.Up, wh1.Down:
			switch currentItem.Item() {
			case RobotItem:
				nextLeft := nextItem.PositionLeft()
				wh.Move(&nextLeft, direction, level)
			case WallItem, BoxItem:
				step := 1
				if direction == wh1.Up {
					step = -1
				}
				newLocationLeft := NewLocation(currentItem.PositionLeft().X(), currentItem.PositionLeft().Y()+step)
				if newItemleft, exists := wh.ItemAtPosition(newLocationLeft); exists && newItemleft.Item() == BoxItem {
					partLeft := newItemleft.PositionLeft()
					wh.Move(&partLeft, direction, level)
				}
				newLocationRight := NewLocation(currentItem.PositionRight().X(), currentItem.PositionRight().Y()+step)
				if newItemRight, exists := wh.ItemAtPosition(newLocationRight); exists && newItemRight.Item() == BoxItem {
					partLeft := newItemRight.PositionLeft()
					wh.Move(&partLeft, direction, level)
				}
			}
		}
	}

    if level == 1 && wh.unwindCause == WallItem {
        fmt.Println("**** Hit a Wall ****", level)
        fmt.Println(wh.undoLog)
        wh.Undo()
        return
    }

	switch currentItem.Item() {
	case BoxItem:
		locationToDelete := currentItem.PositionLeft()
		currentItem.SetPosition(*nextPosition)
		wh.items[*nextPosition] = currentItem
        wh.AddUndoItem(&currentItem, direction)
		delete(wh.items, locationToDelete)
		return
	case RobotItem:
		locationToDelete := wh.robot.PositionLeft()
		wh.robot.position = *nextPosition
		wh.items[*nextPosition] = &wh.robot
		delete(wh.items, locationToDelete)
		return
	}
}

func (wh *Warehouse) GoRobotGo() {
	for i := 0; i < wh.robotPath.Length(); i++ {
		ptr := wh.robotPath.NextPointer()
		wh.Move(&wh.robot.position, ptr, 0)
	}
}

func (wh *Warehouse) AddUndoItem(item *Item, pointer wh1.Pointer) {
    wh.undoLog = append(wh.undoLog, UndoItem{item, pointer})
}

func (wh *Warehouse) ClearUndoLog() {
    wh.undoLog = wh.undoLog[:0]
    wh.unwindCause = UnusedItem
}

func (wh *Warehouse) Undo() {
    for _, undoItem := range wh.undoLog  {
        step := 1
        if undoItem.direction == wh1.Down {
            step = -1
        }
        locationToDelete := (*undoItem.item).PositionLeft()
        originalPosition := NewLocation(locationToDelete.x, locationToDelete.y + step)
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
