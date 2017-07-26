package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	require := require.New(t)
	var input []int
	var target int
	var result int

	input = []int{1, 4, 6, 12, 30, 68}
	target = 5

	result = BinarySearchGreaterThan(input, 0, len(input)-1, target)
	require.Equal(2, result)

	result = BinarySearchLessThan(input, 0, len(input)-1, target)
	require.Equal(1, result)

	result = BinarySearchGreaterThan(input, 2, len(input)-1, target)
	require.Equal(2, result)

	result = BinarySearchLessThan(input, 2, len(input)-1, target)
	require.Equal(1, result)
}
