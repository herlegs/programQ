package easy

func singleNumber(nums []int) int {
	r := 0
	for _, n := range nums {
		r ^= n
	}
	return r
}
