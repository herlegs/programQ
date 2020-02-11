package medium

import (
	"container/heap"
)

func openLock(deadends []string, target string) int {
	start := "0000"
	deadMap := map[string]bool{}
	for _, d := range deadends {
		deadMap[d] = true
	}
	if deadMap[start] {
		return -1
	}
	visited := map[string]bool{}
	pq := NewPriorityQueue()
	pq.Push(&Element{State: start, Step: 0, Total: steps(start, target)})
	visited[start] = true
	for pq.Len() > 0 {
		node := pq.Pop()
		nbrs := neighbors(node.State)
		currStep := node.Step + 1
		for _, n := range nbrs {
			if n == target {
				return currStep
			}
			if !visited[n] && !deadMap[n] {
				pq.Push(&Element{State: n, Step: currStep, Total: currStep + steps(n, target)})
			}
		}
	}
	return -1
}

func steps(a, b string) int {
	s := 0
	for i := 0; i < 4; i++ {
		s += minTurn(a[i], b[i])
	}
	return s
}

func minTurn(a, b uint8) int {
	dist := abs(int(a) - int(b))
	if dist > 5 {
		return 10 - dist
	}
	return dist
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func neighbors(s string) []string {
	return []string{
		string([]byte{(s[0]-'0'+11)%10 + '0', s[1], s[2], s[3]}),
		string([]byte{(s[0]-'0'+9)%10 + '0', s[1], s[2], s[3]}),
		string([]byte{s[0], (s[1]-'0'+11)%10 + '0', s[2], s[3]}),
		string([]byte{s[0], (s[1]-'0'+9)%10 + '0', s[2], s[3]}),
		string([]byte{s[0], s[1], (s[2]-'0'+11)%10 + '0', s[3]}),
		string([]byte{s[0], s[1], (s[2]-'0'+9)%10 + '0', s[3]}),
		string([]byte{s[0], s[1], s[2], (s[3]-'0'+11)%10 + '0'}),
		string([]byte{s[0], s[1], s[2], (s[3]-'0'+9)%10 + '0'}),
	}
}

type Element struct {
	State string
	//steps already walked
	Step int
	//total dist to the target (steps + remaining estimation)
	Total int
}

type PQ []*Element

func (q *PQ) Len() int {
	return len(*q)
}

func (q *PQ) Less(i, j int) bool {
	return (*q)[i].Total < (*q)[j].Total
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
