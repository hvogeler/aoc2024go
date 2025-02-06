package maze2p2

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test_FromString(t *testing.T) {
	t.Run("Example Data1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		// fmt.Println(data)
		m := MazeFromStr(data)
		// fmt.Println(m)

		m.FindPath()
		fmt.Printf("Cost Part1: %d\n", m.Score())
		if m.Score() != 7036 {
			t.Errorf("Expected low score 7036, got %d", m.Score())
		}

		p := []*NodeTile{}
		m.WalkShortestPaths(m.finishTile, p)
		for _, path := range m.ShortestPaths() {
			fmt.Println(m.PrintPath(path))
		}

		fmt.Printf("Number of visited tiles Part 2: %d\n", m.CountAllVisitedTiles())
		if m.CountAllVisitedTiles() != 45 {
			t.Errorf("Expected tiles 45, got %d", m.CountAllVisitedTiles())
		}
	})

	t.Run("Example Data2", func(t *testing.T) {
		data := ReadData("../example2.dat")
		fmt.Println(data)
		m := MazeFromStr(data)
		fmt.Println(m)

		m.FindPath()
		fmt.Printf("Cost: %d\n", m.Score())
		if m.Score() != 11048 {
			t.Errorf("Expected low score 11048, got %d", m.Score())
		}

		p := []*NodeTile{}
		m.WalkShortestPaths(m.finishTile, p)
		for _, path := range m.ShortestPaths() {
			fmt.Println(m.PrintPath(path))
		}


		fmt.Printf("Number of visited tiles Part 2: %d\n", m.CountAllVisitedTiles())
		if m.CountAllVisitedTiles() != 64 {
			t.Errorf("Expected tiles 64, got %d", m.CountAllVisitedTiles())
		}
	})

	t.Run("Example Data5", func(t *testing.T) {
		data := ReadData("../example5.dat")
		fmt.Println(data)
		m := MazeFromStr(data)
		fmt.Println(m)

		m.FindPath()
		fmt.Printf("Cost: %d\n", m.Score())
		if m.Score() != 5078 {
			t.Errorf("Expected low score 5078, got %d", m.Score())
		}

		p := []*NodeTile{}
		m.WalkShortestPaths(m.finishTile, p)
		for _, path := range m.ShortestPaths() {
			fmt.Println(m.PrintPath(path))
		}

		fmt.Printf("Number of visited tiles Part 2: %d\n", m.CountAllVisitedTiles())
		if m.CountAllVisitedTiles() != 413 {
			t.Errorf("Expected tiles 413, got %d", m.CountAllVisitedTiles())
		}
	})

	t.Run("Example Data6", func(t *testing.T) {
		data := ReadData("../example6.dat")
		fmt.Println(data)
		m := MazeFromStr(data)
		fmt.Println(m)

		m.FindPath()
		fmt.Printf("Cost: %d\n", m.Score())
		if m.Score() != 21110 {
			t.Errorf("Expected low score 21110, got %d", m.Score())
		}

		p := []*NodeTile{}
		m.WalkShortestPaths(m.finishTile, p)
		// for _, path := range m.ShortestPaths() {
		// 	fmt.Println(m.PrintPath(path))
		// }

		fmt.Printf("Number of visited tiles Part 2: %d\n", m.CountAllVisitedTiles())
		if m.CountAllVisitedTiles() != 264 {
			t.Errorf("Expected tiles 264, got %d", m.CountAllVisitedTiles())
		}
	})

	t.Run("Example Data9", func(t *testing.T) {
		data := ReadData("../example9.dat")
		fmt.Println(data)
		m := MazeFromStr(data)
		fmt.Println(m)

		m.FindPath()
		fmt.Printf("Cost: %d\n", m.Score())
		if m.Score() != 6027 {
			t.Errorf("Expected low score 6027, got %d", m.Score())
		}

		p := []*NodeTile{}
		m.WalkShortestPaths(m.finishTile, p)
		for i, path := range m.ShortestPaths() {
            fmt.Println(i)
			fmt.Println(m.PrintPath(path))
            fmt.Println()
		}

		fmt.Printf("Number of visited tiles Part 2: %d\n", m.CountAllVisitedTiles())
		if m.CountAllVisitedTiles() != 264 {
			t.Errorf("Expected tiles 264, got %d", m.CountAllVisitedTiles())
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
		nt1 := NewNodeTile(1, 1, Undefined)
		nt2 := NewNodeTile(2, 2, Undefined)
		nt3 := NewNodeTile(3, 3, Undefined)
		nt4 := NewNodeTile(4, 5, Undefined)
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
