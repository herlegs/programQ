package hard

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTakeCoin(t *testing.T) {
	require.Equal(t, true, takeCoin(1))
	require.Equal(t, true, takeCoin(2))
	require.Equal(t, false, takeCoin(3))
	require.Equal(t, true, takeCoin(4))
	require.Equal(t, true, takeCoin(5))
	require.Equal(t, false, takeCoin(6))
	require.Equal(t, true, takeCoin(7))
}

func TestTakeCoinLeft(t *testing.T) {
	require.Equal(t, -1, takeCoinLeft([]int{-1}))
	require.Equal(t, 1, takeCoinLeft([]int{1}))
	require.Equal(t, 2, takeCoinLeft([]int{-1, 3}))
	require.Equal(t, 6, takeCoinLeft([]int{-1, 3, 4, 5}))
	require.Equal(t, 7, takeCoinLeft([]int{6, -1, -3, -1, 5}))
}

func TestTakeCoinSide(t *testing.T) {
	require.Equal(t, 3, takeCoinSide([]int{1, 5, 2}))
}
