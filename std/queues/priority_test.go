package queues

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	ExamplePriorityQueue()
}

func ExamplePriorityQueue() {
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear":   4,
		"tt":     1,
		"orange": 5,
	}
	pq := NewPriorityQueue[string](len(items))
	for value, priority := range items {
		pq.Push(NewItem(priority, value))
	}
	for pq.Len() > 0 {
		item := pq.Pop()
		fmt.Printf("%d:%s ", item.priority, item.value)
	}
}
