package others

import (
	"fmt"
	"testing"
)

func TestFindDroppingBricks(t *testing.T) {
	m := [][]int{
		{1, 0, 1, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	res := FindDroppingBricks(m, 2, 0) - 1
	for _, r := range m {
		fmt.Printf("%v\n", r)
	}

	fmt.Printf("%v\n", res)
}
