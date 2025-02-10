package memspace

import (
	l "container/list"

)

type Queue struct {
	data *l.List
}

func NewQueue() Queue {
	return Queue{
		data: l.New(),
	}
}

func (q *Queue) Enq(memLoc *MemLocation) {
	q.data.PushBack(memLoc)
}


func (q *Queue) Deq() *MemLocation {
	memLocElement := q.data.Front()
	if memLocElement == nil {
		return nil
	}
	q.data.Remove(memLocElement)
	memLoc, ok := memLocElement.Value.(*MemLocation)
	if !ok {
		panic("Invalid type in queue")
	}
	return memLoc
}