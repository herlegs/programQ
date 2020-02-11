package medium

import (
	"fmt"
	"testing"
)

func TestOpenLock(t *testing.T) {
	r := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	fmt.Printf("res:%v\n", r)
}

func TestName(t *testing.T) {
	fmt.Printf("%v\n", steps("0002", "1003"))
}
