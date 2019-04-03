package util

//return root
func find(union map[int]int, curr int) int {
	if _, exist := union[curr]; !exist {
		union[curr] = curr
		return curr
	}
	if union[curr] != curr {
		union[curr] = find(union, union[curr])
	}
	return union[curr]
}

//return removedCount
func union(u map[int]int, a, b int) int {
	aRoot := find(u, a)
	bRoot := find(u, b)
	if aRoot == bRoot {
		return 0
	}
	u[aRoot] = bRoot
	return 1
}
