package util

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	queue.Push("abc")
	queue.Push("edf")
	require.True(t, queue.Size() == 2)
	require.Equal(t, "abc", queue.Pop())
	require.Equal(t, "edf", queue.Pop())
	require.True(t, queue.IsEmpty())
	require.Equal(t, nil, queue.Pop())
}

func TestNewStack(t *testing.T) {
	stack := NewStack()
	stack.Push("abc")
	stack.Push("ggg")
	require.True(t, stack.Size() == 2)
	require.Equal(t, "ggg", stack.Pop())
	require.Equal(t, "abc", stack.Pop())
	require.True(t, stack.IsEmpty())
	require.Equal(t, nil, stack.Pop())
}