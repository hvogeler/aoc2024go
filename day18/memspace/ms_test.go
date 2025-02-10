package memspace

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	t.Run("Example Data1", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		memSpace := MemSpaceFromStr(data, 7, 7, 12)
		fmt.Println(memSpace)
	})
}

func Test_queue(t *testing.T) {
	t.Run("Queue Test", func(t *testing.T) {
		data := ReadData("../example1.dat")
		fmt.Println(data)
		memSpace := MemSpaceFromStr(data, 7, 7, 12)
		q := NewQueue()
		q.Enq(memSpace.memLocations[NewLocation(0, 0)])
		q.Enq(memSpace.memLocations[NewLocation(1, 5)])
		q.Enq(memSpace.memLocations[NewLocation(6, 6)])

		for e := q.data.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}

		mem0 := q.Deq()
		fmt.Println(mem0.pos, mem0.memType)
        if mem0.memType != Start {
            t.Errorf("Expected Start, got %s", mem0.memType)
        }
		mem1 := q.Deq()
		fmt.Println(mem1.pos, mem1.memType)
        if mem1.memType != Corrupt {
            t.Errorf("Expected Corrupt, got %s", mem1.memType)
        }
		mem2 := q.Deq()
		fmt.Println(mem2.pos, mem2.memType)
        if mem2.memType != Exit {
            t.Errorf("Expected Exit, got %s", mem2.memType)
        }
	})
}
