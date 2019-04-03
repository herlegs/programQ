package hard

import "fmt"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if k > n {
		profits := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				profits += prices[i] - prices[i-1]
			}
		}
		return profits
	}
	profits := make([][]int, 2)
	for i := 0; i < n; i++ {
		profits[i] = make([]int, n)
	}
	for x := 1; x <= k; x++ {
		for i := 1; i < n; i++ {
			maxP := 0
			for j := i - 1; j >= 0; j-- {
				profit := maxInt(0, prices[i]-prices[j]) + profits[j][(x-1)%2]
				maxP = maxInt(maxP, profit)
			}
			profits[i][x%2] = maxInt(maxP, profits[i-1][x%2])
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%v%v\n", prices[i], profits[i])
	}

	maxP := 0
	for i := 0; i < n; i++ {
		maxP = maxInt(maxP, profits[i][k%2])
	}

	return maxP
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
