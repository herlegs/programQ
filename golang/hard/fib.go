package hard

func FibIterate(n int) int {
	if n <= 1 {
		return 1
	}
	prev1, prev2 := 1, 1
	for n > 1 {
		n--
		prev1, prev2 = prev2, prev1+prev2
	}
	return prev2
}

func Fib(n int) int {
	//[fn, fn-1] = matrix[1,1,1,0] * [fn-1, fn-2]

	//[fn, fn-1] = matrix[1,1,1,0]^n-1 * [f1, f0] = matrix[1,1,1,0]^n-1 * [1, 1]

	m := matrixPow([4]int{1, 1, 1, 0}, n-1)
	return m[0] + m[1]
}

func matrixPow(m [4]int, pow int) [4]int {
	result := [4]int{1, 0, 0, 1}
	base := m
	for pow > 0 {
		shouldMul := pow&1 == 1
		if shouldMul {
			result = matrixMul(result, base)
		}
		base = matrixMul(base, base)
		pow = pow >> 1
	}
	return result
}

func matrixMul(m, n [4]int) [4]int {
	r := [4]int{
		m[0]*n[0] + m[1]*n[2],
		m[0]*n[1] + m[1]*n[3],
		m[2]*n[0] + m[3]*n[2],
		m[2]*n[1] + m[3]*n[3],
	}
	return r
}
