package util

import (
	"fmt"
	"testing"
)

func TestUnionFind(t *testing.T) {
	fmt.Printf("what\n")
	matrix := [][]int{
		{1, 1, 0, 1},
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{0, 1, 1, 0},
	}
	n, m := len(matrix), len(matrix[0])
	u := map[int]int{}
	added, removed := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			added++
			if j-1 >= 0 && matrix[i][j-1] == 1 {
				r := union(u, i*m+j-1, i*m+j)
				removed += r
			}
			if i-1 >= 0 && matrix[i-1][j] == 1 {
				r := union(u, (i-1)*m+j, i*m+j)
				removed += r
			}
		}
	}
	fmt.Printf("total:%v\n", added-removed)
}
