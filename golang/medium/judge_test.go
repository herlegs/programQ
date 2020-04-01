package medium

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	matrix := [][]int{{0, 0, 0, 0, 1, 1, 1, 1, 0}, {0, 1, 1, 0, 0, 0, 0, 1, 0}, {0, 0, 1, 0, 0, 0, 0, 0, 0}, {1, 1, 0, 0, 1, 0, 0, 1, 1}, {0, 0, 1, 1, 1, 0, 1, 0, 1}, {0, 1, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 1, 0, 0, 0}, {0, 1, 0, 1, 1, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 1, 0, 1, 0}}
	r := shortestPathBinaryMatrix(matrix)
	fmt.Printf("res:%v\n", r)
}

func TestName(t *testing.T) {
	a := "dfdfd"
	for _, b := range a {
		fmt.Printf("%T %v\n", b, b)
	}
}
