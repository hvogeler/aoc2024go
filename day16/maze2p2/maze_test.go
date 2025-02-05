package maze2p2

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

		m.FindPath()
		p := []*NodeTile{}
		m.WalkShortestPaths(m.finishTile, p)
		fmt.Println(m.PrintPath(m.ShortestPaths()[2]))
		// fmt.Println(m)
		// for _, tile := range m.tiles {
		// 	if nt, ok := tile.(*NodeTile); ok {
		// 		if len(nt.preTile) > 1 {
		// 			fmt.Println(nt.pos)
		// 		}
		// 	}
		// }
		// fmt.Println(m.PrintPath())
		fmt.Printf("Cost: %d\n", m.Score())
		if m.Score() != 7036 {
			t.Errorf("Expected low score 7036, got %d", m.Score())
		}
	})

	t.Run("Example Data2", func(t *testing.T) {
		data := ReadData("../example2.dat")
		fmt.Println(data)
		m := MazeFromStr(data)
		fmt.Println(m)

		m.FindPath()
		fmt.Println(m)
		// fmt.Println(m.PrintPath())
		fmt.Printf("Cost: %d\n", m.Score())
		if m.Score() != 11048 {
			t.Errorf("Expected low score 11048, got %d", m.Score())
		}
	})

	t.Run("Example Data3", func(t *testing.T) {
		data := ReadData("../example3.dat")
		fmt.Println(data)
		m := MazeFromStr(data)
		fmt.Println(m)

		m.FindPath()
		fmt.Println(m)
		// fmt.Println(m.PrintPath())
		fmt.Printf("Cost: %d\n", m.Score())
		if m.Score() != 21148 {
			t.Errorf("Expected low score 21148, got %d", m.Score())
		}
	})

	t.Run("Example Data4", func(t *testing.T) {
		data := ReadData("../example4.dat")
		fmt.Println(data)
		m := MazeFromStr(data)
		fmt.Println(m)

		m.FindPath()
		fmt.Println(m)
		// fmt.Println(m.PrintPath())
		fmt.Printf("Cost: %d\n", m.Score())
		if m.Score() != 4013 {
			t.Errorf("Expected low score 4013, got %d", m.Score())
		}
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
