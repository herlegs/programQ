package util

type BSType int

const (
	LESS BSType = iota
	GREATER
)

// return target value's matched element's index within start ~ end (inclusive) index of array
// if no match, depending BSType, return element's index with less or greater value than target
func BinarySearch(array []int, start, end, target int, bsType BSType) int {
	/*for start <= end {
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
	}*/
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
	}
	return start
}

func BinarySearchLessThan(array []int, start, end, target int) int {
	return BinarySearch(array, start, end, target, LESS)
}

func BinarySearchGreaterThan(array []int, start, end, target int) int {
	return BinarySearch(array, start, end, target, GREATER)
}

//first element bigger than target (so all element on left side is smaller or equal to target)
func BSFirstBigger(array []int, start, end, target int) int {
	for start <= end {
		mid := start + (end-start)/2
		if array[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

func BSLastSmaller(array []int, start, end, target int) int {
	for start <= end {
		mid := start + (end-start)/2
		if array[mid] >= target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return end
}