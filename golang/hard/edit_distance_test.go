package hard

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEditDistance(t *testing.T) {
	data := []struct {
		n   string
		m   string
		dis int
	}{
		{"a", "b", 1},
		{"abc", "acb", 2},
		{"ag", "b", 2},
		{"cag", "bab", 2},
	}

	for _, d := range data {
		res := EditDistanceIterativeMemoryOptimize(d.n, d.m)
		t.Logf("[%v][%v][%v]\n", d.n, d.m, res)
		require.Equal(t, d.dis, res)
	}

}

/*
EditDistance("Where r u mad", "wher3 are you mom") using brute force:
8656812799 ns

using recursive dp
8564 ns

using iterative dp
4803 ns

*/
func BenchmarkEditDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EditDistanceIterative("Where r u mad", "wher3 are you mom")
	}
}
