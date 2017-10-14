package util

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestTreeInorderRecursive(t *testing.T) {
	tree := Deserialize("1,2,#,4,5,#,#,#,6")
	res := TreeInorderRecursive(tree)
	require.Equal(t, []int64{4,2,5,6,1}, res)
}

func TestTreeInorderIterative(t *testing.T) {
	tree := Deserialize("1,2,#,4,5,#,6,#,7")
	res := TreeInorderIterative(tree)
	require.Equal(t, []int64{4,6,2,5,7,1}, res)

	tree = Deserialize("6")
	res = TreeInorderIterative(tree)
	require.Equal(t, []int64{6}, res)
}

func BenchmarkTreeInorderRecursive(b *testing.B) {
	for i := 0; i < b.N; i++{
		tree := Deserialize("1,2,#,4,5,#,6,#,7")
		TreeInorderRecursive(tree)
	}
}

func BenchmarkTreeInorderIterative(b *testing.B) {
	for i := 0; i < b.N; i++{
		tree := Deserialize("1,2,#,4,5,#,6,#,7")
		TreeInorderIterative(tree)
	}
}

func TestTreeInorderIterator(t *testing.T){
	tree := Deserialize("1,2,#,4,5,#,6,#,7")
	inorderT := TreeInorderIterator(tree)
	result := []int64{}
	for inorderT.HasNext() {
		result = append(result, inorderT.Next().Val)
	}
	require.Equal(t, []int64{4,6,2,5,7,1}, result)
}

func TestTreeReverseInorderIterator(t *testing.T){
	tree := Deserialize("1,2,#,4,5,#,6,#,7")
	reverse := TreeReverseInorderIterator(tree)
	result := []int64{}
	for reverse.HasNext() {
		result = append(result, reverse.Next().Val)
	}
	require.Equal(t, []int64{1,7,5,2,6,4}, result)
}

