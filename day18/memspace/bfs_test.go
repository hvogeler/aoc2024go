package memspace

import (
	"fmt"
	"testing"
)

func Test_bfs(t *testing.T) {
	t.Run("Bfs1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		ms := MemSpaceFromStr(data, 7, 7, 12)
		fmt.Println(ms)
		ms.BfsWalk()
		fmt.Printf("Steps for shortest path: %d\n", ms.ExitNode().pathLen)
		if ms.ExitNode().pathLen != 22 {
			t.Errorf("Expected Pathlen 22, got %d", ms.ExitNode().pathLen)
		}
	})
}

func Test_neighbors(t *testing.T) {
	t.Run("Neighbors", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		ms := MemSpaceFromStr(data, 7, 7, 12)
		fmt.Println(ms)
		n1 := ms.Neighbors(ms.StartNode())
		for _, n := range n1 {
			fmt.Println(*n)
		}
		if len(n1) != 2 {
			t.Errorf("Incorrect Neighbors")
		}
		if n1[0].pos != NewLocation(1, 0) {
			t.Errorf("Incorrect Neighbor")
		}
		if n1[1].pos != NewLocation(0, 1) {
			t.Errorf("Incorrect Neighbor")
		}
		fmt.Println()
		n2 := ms.Neighbors(ms.GetAtPos(4, 1))
		for _, n := range n2 {
			fmt.Println(*n)
		}
		if len(n2) != 4 {
			t.Errorf("Incorrect Neighbors")
		}
		if n2[0].pos != NewLocation(3, 1) {
			t.Errorf("Incorrect Neighbor")
		}
		if n2[3].pos != NewLocation(4, 0) {
			t.Errorf("Incorrect Neighbor")
		}
		if n2[1].memType != Corrupt || n2[2].memType != Corrupt {
			t.Errorf("Incorrect Neighbor Type")
		}
})
}
