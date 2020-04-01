package hard

const (
	rows = 1000000
	cols = 1000000
)

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	stones := map[int]bool{}
	for _, b := range blocked {
		key := b[0]*rows + b[1]
		stones[key] = true
	}
	checkSize := maxSize(len(blocked))
	reached, sourceNotBlocked := checkBlock(stones, checkSize, source, target)
	if reached {
		return true
	}
	reached, targetNotBlocked := checkBlock(stones, checkSize, target, source)
	if reached {
		return true
	}

	return sourceNotBlocked && targetNotBlocked
}

//return reached, not_blocked
func checkBlock(stones map[int]bool, checkSize int, source []int, target []int) (bool, bool) {
	visited := map[int]bool{}
	pq := NewQueue()
	start := &Element{X: source[0], Y: source[1], Key: source[0]*rows + source[1]}
	end := &Element{X: target[0], Y: target[1], Key: target[0]*rows + target[1]}
	pq.Push(start)
	visited[start.Key] = true

	for pq.Size() > 0 {
		if len(visited) > checkSize {
			return false, true
		}
		node := pq.Pop().(*Element)
		nbrs := neighbors(node)
		for _, nbr := range nbrs {
			if dist(nbr, end) == 0 {
				return true, true
			}
			if !visited[nbr.Key] && !stones[nbr.Key] {
				visited[nbr.Key] = true
				pq.Push(nbr)
			}
		}
	}
	return false, false
}

func dist(a, b *Element) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func maxSize(l int) int {
	return l * l / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func neighbors(cur *Element) []*Element {
	nbrs := []*Element{}
	for _, dir := range dirs {
		x, y := cur.X+dir[0], cur.Y+dir[1]
		if x < 0 || x >= rows || y < 0 || y >= cols {
			continue
		}
		e := &Element{X: x, Y: y, Key: x*rows + y}
		nbrs = append(nbrs, e)
	}
	return nbrs
}

type Element struct {
	X, Y int
	Key  int
}

type Queue interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
	IsEmpty() bool
}

func NewQueue() Queue {
	return &queueImpl{
		nil,
		nil,
		0,
	}
}

//Single linked list node
type SNode struct {
	Val  interface{}
	Next *SNode
}

// Not thread safe
type queueImpl struct {
	head *SNode
	tail *SNode
	size int
}

func (q *queueImpl) Push(v interface{}) {
	if q.size == 0 {
		q.head = &SNode{Val: v}
		q.tail = q.head
		q.size++
		return
	}
	added := &SNode{Val: v}
	q.tail.Next = added
	q.tail = added
	q.size++
}

func (q *queueImpl) Pop() interface{} {
	if q.size == 0 {
		return nil
	}
	item := q.head.Val
	q.head = q.head.Next
	q.size--
	return item
}

func (q *queueImpl) Peek() interface{} {
	if q.head != nil {
		return q.head.Val
	}
	return nil
}

func (q *queueImpl) Size() int {
	return q.size
}

func (q *queueImpl) IsEmpty() bool {
	return q.size == 0
}
