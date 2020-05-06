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

func TestCanWinLose(t *testing.T) {
	a := "++++-++++-++++-++++"
	fmt.Printf("can win:%v:%v\n", a, canWin(a))
	fmt.Printf("can lose:%v:%v\n", a, canLose(a))
}
