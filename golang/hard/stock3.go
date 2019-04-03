package hard

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	left, right := make([]int, n), make([]int, n)
	min, max := prices[0], prices[n-1]
	for i, j := 1, n-2; i < n && j >= 0; {
		left[i] = maxInt(prices[i]-min, left[i-1])
		if prices[i] < min {
			min = prices[i]
		}
		right[j] = maxInt(max-prices[j], right[j+1])
		if prices[j] > max {
			max = prices[j]
		}
		i++
		j--
	}
	maxProfit := 0
	for k := 0; k < n; k++ {
		if left[k]+right[k] > maxProfit {
			maxProfit = left[k] + right[k]
		}
	}
	return maxProfit
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
