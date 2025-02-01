package maze2

import (
	"fmt"
	"strings"
)

type NodeHeap []NodeTile

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(NodeTile))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h NodeHeap) String() string {
	var s strings.Builder
	for _, nt := range h {
		s.WriteString(fmt.Sprintf("Tile %s, Cost: %d\n", nt.pos, nt.cost))
	}
	return s.String()
}
