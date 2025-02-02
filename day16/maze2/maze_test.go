package maze2

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test_FromString(t *testing.T) {
	t.Run("Example Data1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		m := MazeFromStr(data)
		fmt.Println(m)

		x := m.tiles[NewPosition(1,3)].(*NodeTile)
		x.cost = 42
		fmt.Println(m.tiles[NewPosition(1,3)].(*NodeTile).cost)
		if m.tiles[NewPosition(1,3)].(*NodeTile).cost != 42 {
			t.Errorf("Expected 42")
		}

		m.FindPath()
		fmt.Println(m)
	})
}

func Test_PriorityQueue(t *testing.T) {
	t.Run("Prioqueue1", func(t *testing.T) {
		nt1 := NewNodeTile(1, 1)
		nt2 := NewNodeTile(2, 2)
		nt3 := NewNodeTile(3, 3)
		nt4 := NewNodeTile(4, 5)
		nt1.cost = 5
		nt2.cost = 1
		nt3.cost = 2
		nt4.cost = 4
		h := &NodeHeap{nt1, nt2, nt3}
		fmt.Println(h)
		heap.Init(h)
		fmt.Println(h)
		heap.Push(h, nt4)
		fmt.Println(h)

		l := (*h).Len()
		for i := 0; i < l; i++ {
			topNt := heap.Pop(h).(*NodeTile)
			fmt.Printf("Top Tile %s, Cost: %d\n", topNt.pos, topNt.cost)
		}
	})
}
