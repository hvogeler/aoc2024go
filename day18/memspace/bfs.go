package memspace

func (ms *MemSpace) BfsWalk() {
	q := NewQueue()
	startNode := ms.StartNode()
	q.Enq(startNode)
	// path := []MemLocation{*startNode}
	for stepCount := 0; q.data.Len() > 0; stepCount++ {
		node := q.Deq()
		neighbors := ms.Neighbors(node)
		for i := 0; i < len(neighbors); i++ {
			neighbor := neighbors[i]
			if neighbor.memType == Corrupt || neighbor.IsVisited() {
				continue
			}
			neighbor.pathLen = node.pathLen+1
			neighbor.prev = node
			neighbor.memType = Visited
			q.Enq(neighbor)
		}
	}
}

func (ms *MemSpace) Neighbors(node *MemLocation) []*MemLocation {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, 1, -1}
	memLocs := []*MemLocation{}
	for i :=0; i < 4; i++ {
		if node.pos.x + dx[i] < 0 || node.pos.x + dx[i] >= ms.dimensions.dimX {
			continue
		}
		if node.pos.y + dy[i] < 0 || node.pos.y + dy[i] >= ms.dimensions.dimY {
			continue
		}
		pos := NewLocation(node.pos.x + dx[i], node.pos.y + dy[i])
		memLoc := ms.GetAt(pos)
		memLocs = append(memLocs, memLoc)
	}
	return memLocs
}