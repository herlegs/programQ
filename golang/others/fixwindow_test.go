package others

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFixWindow(t *testing.T) {
	r := fixWindowMax([]int{1, 1, 2, 2, 3, 4, 5, 6}, 3)
	require.Equal(t, 6, r)

	r = fixWindowMax([]int{1, 1, 2, 2, 3, 4, 5, 5, 5, 6}, 3)
	require.Equal(t, 7, r)

	r = fixWindowMax([]int{1}, 3)
	require.Equal(t, 1, r)
}

func TestTimeBucket(t *testing.T) {
	n, f := time.Now().Zone()
	fmt.Printf("%v,%v\n", n, f)
}
