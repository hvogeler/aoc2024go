package warehouse2

import (
	wh1 "day15/warehouse"
	"fmt"
	"strings"
)


type UndoItem struct {
	item      *Item
	direction wh1.Pointer
}

func (u UndoItem) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("%s at %s", (*u.item).Item(), (*u.item).PositionLeft()))
	return s.String()
}