package util

import (
	"fmt"
	"testing"
)

func TestPQ(t *testing.T) {
	q := NewPriorityQueue()
	q.Push(&Element{Val: 2})
	q.Push(&Element{Val: 5})
	q.Push(&Element{Val: 3})
	q.Push(&Element{Val: 4})
	q.Push(&Element{Val: 1})

	var r []int
	for q.Len() > 0 {
		r = append(r, q.Pop().Val)
	}

	fmt.Printf("%v\n", r)
}
