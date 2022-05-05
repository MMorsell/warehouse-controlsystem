package pathfinding

import (
	"container/heap"
	"testing"
)

func TestPriority(t *testing.T) {
	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &item{Position{1, 25, 1}, 12})
	heap.Push(pq, &item{Position{5, 0, 1}, 5})
	heap.Push(pq, &item{Position{0, 0, 1}, 214})
	heap.Push(pq, &item{Position{-1, -5, 1}, 2})

	priorityOrder := []int32{2, 5, 12, 214}

	for _, want := range priorityOrder {
		got := heap.Pop(pq).(*item).priority
		if got != want {
			t.Errorf("Expected %d, but got %d", want, got)
		}
	}
}
