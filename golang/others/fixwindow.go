package others

/*
give sorted list, and a window size.
Find the max size of the sub array, in which the diff of max and min value does not exceed window size
*/
func fixWindowMax(arr []int, window int) int {
	n := len(arr)
	if n == 0 || window <= 0 {
		return 0
	}
	maxCount := 0
	left, right := 0, 0
	for left <= right && right < n {
		for left <= right && right < n && arr[right]-arr[left] > window {
			left++
		}
		maxCount = max(maxCount, right-left+1)
		right++
	}
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
