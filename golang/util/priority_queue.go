package util

import "container/heap"

type Element struct {
	Val int
}

type PQ []*Element

func (q *PQ) Len() int {
	return len(*q)
}

func (q *PQ) Less(i, j int) bool {
	return (*q)[i].Val < (*q)[j].Val
}

func (q *PQ) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *PQ) Push(v interface{}) {
	*q = append(*q, v.(*Element))
}

func (q *PQ) Pop() interface{} {
	last := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return last
}

type PriorityQueue interface {
	Len() int
	Push(ele *Element)
	Pop() *Element
}

type priorityQueueImpl struct {
	*PQ
}

func NewPriorityQueue() PriorityQueue {
	return &priorityQueueImpl{
		&PQ{},
	}
}

func (pq *priorityQueueImpl) Len() int {
	return pq.PQ.Len()
}

func (pq *priorityQueueImpl) Push(e *Element) {
	heap.Push(pq.PQ, e)
}

func (pq *priorityQueueImpl) Pop() *Element {
	return heap.Pop(pq.PQ).(*Element)
}
