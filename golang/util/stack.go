package util

import (
	"container/list"
)

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
	IsEmpty() bool
}

func NewStack() Stack {
	return &stackImpl{
		l: list.New(),
	}
}

type stackImpl struct {
	l *list.List
}

func (s *stackImpl) Push(v interface{}) {
	s.l.PushBack(v)
}

func (s *stackImpl) Pop() interface{} {
	back := s.l.Back()
	if back == nil {
		return nil
	}
	s.l.Remove(back)
	return back.Value
}

func (s *stackImpl) Peek() interface{} {
	back := s.l.Back()
	if back == nil {
		return nil
	}
	return back.Value
}

func (s *stackImpl) Size() int {
	return s.l.Len()
}

func (s *stackImpl) IsEmpty() bool {
	return s.l.Len() == 0
}
