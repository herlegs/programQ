package util

import (
	"fmt"
	"testing"
)

func TestPQ(t *testing.T) {
	q := NewPriorityQueue()
	q.Push(&Element{Val: 2})
	q.Push(&Element{Val: 5})
	q.Push(&Element{Val: 3})
	q.Push(&Element{Val: 4})
	q.Push(&Element{Val: 1})

	var r []int
	for q.Len() > 0 {
		r = append(r, q.Pop().Val)
	}

	fmt.Printf("%v\n", r)
}

func getDigits(s string) (rune, rune, rune, rune) {
	return rune(s[0] - '0'), rune(s[1] - '0'), rune(s[2] - '0'), rune(s[3] - '0')
}

func toStr(a, b, c, d rune) string {
	a, b, c, d = (a+10)%10, (b+10)%10, (c+10)%10, (d+10)%10
	return string([]rune{'0' + a, '0' + b, '0' + c, '0' + d})
}

func neighbors(s string) []string {
	return []string{
		string([]byte{(s[0]-'0'+11)%10 + '0', s[1], s[2], s[3]}),
		string([]byte{(s[0]-'0'+9)%10 + '0', s[1], s[2], s[3]}),
		string([]byte{s[0], (s[1]-'0'+11)%10 + '0', s[2], s[3]}),
		string([]byte{s[0], (s[1]-'0'+9)%10 + '0', s[2], s[3]}),
		string([]byte{s[0], s[1], (s[2]-'0'+11)%10 + '0', s[3]}),
		string([]byte{s[0], s[1], (s[2]-'0'+9)%10 + '0', s[3]}),
		string([]byte{s[0], s[1], s[2], (s[3]-'0'+11)%10 + '0'}),
		string([]byte{s[0], s[1], s[2], (s[3]-'0'+9)%10 + '0'}),
	}
}

func TestName(t *testing.T) {
	fmt.Printf("%v\n", neighbors("9000"))
}
