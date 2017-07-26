package hard

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{2}
	nums2 := []int{1, 3, 4, 5}

	m := findMedianSortedArrays(nums1, nums2)
	require.Equal(t, 3.0, m)

	nums1 = []int{1, 3}
	nums2 = []int{2}

	m = findMedianSortedArrays(nums1, nums2)
	require.Equal(t, 2.0, m)
}
