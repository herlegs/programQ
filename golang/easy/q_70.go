package easy

func climbStairs(n int) int {
	res := make([]int, 2)
	res[0] = 1
	res[1] = 2
	for i := 2; i < n; i++ {
		res[i%2] = res[(i-1)%2] + res[(i-2)%2]
	}
	return res[(n-1)%2]
}
