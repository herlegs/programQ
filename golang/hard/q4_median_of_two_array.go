package hard

import (
	"math"
)

/*
https://leetcode.com/problems/median-of-two-sorted-arrays/#/description
*/

type BSType int

const (
	LESS BSType = iota
	GREATER
)

type Range struct {
	start int
	end   int
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalCount := len(nums1) + len(nums2)
	leftCount := totalCount / 2

	medArrayId, medIdx := findMedianInRange(nums1, nums2,
		&Range{0, len(nums1) - 1}, &Range{0, len(nums2) - 1},
		leftCount)

	var medArray, theOther []int
	if medArrayId == 1 {
		medArray = nums1
		theOther = nums2
	} else {
		medArray = nums2
		theOther = nums1
	}

	med := medArray[medIdx]

	if totalCount%2 != 0 {
		return float64(med)
	}

	var smaller int = med
	found := BinarySearchLessThan(theOther, 0, len(theOther)-1, med)
	if found > 0 && found < len(theOther) {
		smaller = theOther[found]
	}
	if medIdx > 0 {
		smaller = Min(smaller, medArray[medIdx-1])
	}

	return (float64(med) + float64(smaller)) / 2
}

func findMedianInRange(num1, num2 []int, range1, range2 *Range, leftCount int) (int, int) {
	for range1.start <= range1.end || range2.start <= range2.end {
		ok, arrayId, medIdx := findMedianAndShrink(num1, num2, range1, range2, leftCount, 1)
		if ok {
			return arrayId, medIdx
		}
		ok, arrayId, medIdx = findMedianAndShrink(num2, num1, range2, range1, leftCount, 2)
		if ok {
			return arrayId, medIdx
		}
	}

	return 1, range1.start
}

func findMedianAndShrink(a, b []int, rangeA, rangeB *Range, leftCount int, arrayId int) (bool, int, int) {
	if rangeA.start > rangeA.end {
		return false, arrayId, -1
	}
	mid, found, count := getLeftCount(a, b, rangeA, rangeB)
	if count == leftCount {
		return true, arrayId, mid
	} else if count < leftCount {
		rangeA.start = mid + 1
		rangeB.start = found + 1
	} else {
		rangeA.end = mid - 1
		rangeB.end = found
	}
	return false, arrayId, mid
}

func getLeftCount(num1, num2 []int, range1, range2 *Range) (int, int, int) {
	mid1 := range1.start + (range1.end-range1.start)/2
	found1 := BinarySearchLessThan(num2, range2.start, range2.end, num1[mid1])
	count := mid1 + found1 + 1
	return mid1, found1, count
}

func BinarySearch(array []int, start, end, target int, bsType BSType) int {
	for start <= end {
		mid := start + (end-start)/2

		if array[mid] > target {
			end = mid - 1
		} else if array[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	if bsType == LESS {
		return end
	} else {
		return start
	}
}

func BinarySearchLessThan(array []int, start, end, target int) int {
	return BinarySearch(array, start, end, target, LESS)
}

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
