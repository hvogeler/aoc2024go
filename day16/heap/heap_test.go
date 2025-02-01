package heap

import (
	"fmt"
	"testing"
	"container/heap"
	
)


func Test_FromString(t *testing.T) {
	t.Run("Int Heap", func(t *testing.T) {
		h := &IntHeap{2, 1, 5}
		heap.Init(h)
		heap.Push(h, 3)
		heap.Push(h, 7)
		fmt.Printf("minimum: %d\n", (*h)[0])
		for h.Len() > 0 {
			popped := heap.Pop(h)
			fmt.Printf("%d ", popped)
		}
	
	})
}
