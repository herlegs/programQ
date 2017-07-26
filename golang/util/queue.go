package util

type Queue interface {
	Push(interface{})
	Pop() interface{}
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

func (q *queueImpl) Pop() interface{}{
	if q.size == 0 {
		return nil
	}
	item := q.head.Val
	q.head = q.head.Next
	q.size--
	return item
}

func (q *queueImpl) Size() int {
	return q.size
}

func (q *queueImpl) IsEmpty() bool {
	return q.size == 0
}
