package hard

func takeCoin(n int) bool {
	if n <= 0 {
		return false
	}
	if n <= 2 {
		return true
	}
	dp := make([]bool, 2)
	dp[0] = false
	dp[1] = true
	for i := 2; i <= n; i++ {
		dp[i%2] = !dp[(i-1)%2] || !dp[(i-2)%2]
	}
	return dp[n%2]
}

/*
dp[i] (coins[i] + sum[i+1]-dp[i+1]) coins[i] + coins[i+1] + sum[i+2]-dp[i+2]
*/

// take 1, 2 or 3 coin
func takeCoinLeft(coins []int) int {
	n := len(coins)
	dp := make([]int, 3)
	sum := 0
	for i := n - 1; i >= 0; i-- {
		sum = sum + coins[i]
		otherMin := dp[(i+1)%3]
		if i+2 <= n {
			otherMin = min(otherMin, dp[(i+2)%3])
		}
		if i+3 <= n {
			otherMin = min(otherMin, dp[(i+3)%3])
		}
		dp[i%3] = sum - otherMin
	}
	return dp[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
dp[i,j] = coins[i] + sum[i+1,j]-dp[i+1,j] / coins[j] + sum[i,j-1]-dp[i,j-1]
*/
//take 1 coin from left or right
func takeCoinSide(nums []int) int {
	n := len(nums)
	//sum[i] is for sum to ith element
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	//dp[i][j] is for max score given array index i~j
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = nums[i]
	}
	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			otherMin := min(dp[i][j-1], dp[i+1][j])
			//sum of index i~j
			dp[i][j] = sum[j+1] - sum[i] - otherMin
		}
	}
	return dp[0][n-1]
}

/*
1, 2, 3, 4, 5   10
*/
