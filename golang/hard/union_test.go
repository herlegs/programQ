package hard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnion(t *testing.T) {
	result := calcEquation([][]string{{"a", "b"}, {"e", "f"}, {"b", "e"}}, []float64{3.4, 1.4, 2.3}, [][]string{{"a", "f"}})

	require.EqualValues(t, []float64{10.948}, result)
}

func TestMap(t *testing.T) {
	m := make(map[string]*Parent)
	m["a"] = &Parent{"a", 1}
	m["a"].key = "b"
	m["a"].ratio = 2
	fmt.Printf("%#v\n", m["a"])
}
