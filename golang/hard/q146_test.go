package hard

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLRUCache(t *testing.T) {
	capacity := 1
	cache := Constructor(capacity)
	cache.Put(2, 1)
	require.Equal(t, 1, cache.Get(2))

	cache.Put(3, 2)
	require.Equal(t, -1, cache.Get(2))

	require.Equal(t, 2, cache.Get(3))
}
