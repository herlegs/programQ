package medium

import (
	"container/heap"
)

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	start := &Element{X: 0, Y: 0, Step: 1}
	end := &Element{X: n - 1, Y: m - 1}
	if grid[start.X][start.Y] != 0 || grid[end.X][end.Y] != 0 {
		return -1
	}
	if steps(start, end) == 0 {
		return 1
	}
	start.Total = start.Step + steps(start, end)
	pq := NewPriorityQueue()
	visited := map[int]bool{}
	distMap := map[int]int{}
	pq.Push(start)
	for pq.Len() > 0 {
		node := pq.Pop()
		visited[node.Key] = true
		//fmt.Printf("x:%v,y:%v, %v(%v+%v)\n", node.X, node.Y, node.Total, node.Step, node.Total-node.Step)
		nbrs := neighbors(node, end, n, m)
		for _, nbr := range nbrs {
			if steps(nbr, end) == 0 {
				return nbr.Total
			}
			if visited[nbr.Key] {
				continue
			}
			//old distance
			dist, exist := distMap[nbr.Key]
			if grid[nbr.X][nbr.Y] == 0 && (!exist || nbr.Total < dist) {
				pq.Push(nbr)
				distMap[nbr.Key] = nbr.Total
			}
		}
	}
	return -1
}

func steps(a, b *Element) int {
	return max(abs(a.X-b.X), abs(a.Y-b.Y))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func neighbors(curr, target *Element, n, m int) []*Element {
	nextStep := curr.Step + 1
	res := []*Element{}
	for _, dir := range dirs {
		x, y := curr.X+dir[0], curr.Y+dir[1]
		if x < 0 || x >= n || y < 0 || y >= m {
			continue
		}
		nbr := &Element{X: x, Y: y, Key: x*n + y, Step: nextStep}
		nbr.Total = nbr.Step + steps(nbr, target)
		res = append(res, nbr)
	}
	return res
}

type Element struct {
	X, Y int
	Key  int
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
