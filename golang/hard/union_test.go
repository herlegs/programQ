package hard

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnion(t *testing.T) {
	result := calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2, 3}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}})

	require.EqualValues(t, []float64{6, 0.5, -1, 1, -1}, result)
}
