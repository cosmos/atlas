package crawl_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/atlas/server/crawl"
)

func TestNodePool_Seed(t *testing.T) {
	np := crawl.NewNodePool(10)
	require.Equal(t, 0, np.Size())

	seeds := make([]string, 10)
	for i := range seeds {
		seeds[i] = fmt.Sprintf("127.0.0.%d:26657", i+1)
	}

	np.Seed(seeds)
	require.Equal(t, len(seeds), np.Size())

	np.Seed(seeds)
	require.Equal(t, len(seeds), np.Size())
}

func TestNodePool_RandomNode(t *testing.T) {
	np := crawl.NewNodePool(10)

	seeds := map[string]struct{}{
		"127.0.0.1:26657": {},
		"127.0.0.2:26657": {},
		"127.0.0.3:26657": {},
	}

	rs, ok := np.RandomNode()
	require.False(t, ok)
	require.Empty(t, rs)
	require.NotContains(t, seeds, rs)

	for s := range seeds {
		np.AddNode(s)
	}

	rs, ok = np.RandomNode()
	require.True(t, ok)
	require.Contains(t, seeds, rs)
}

func TestNodePool_AddNode(t *testing.T) {
	np := crawl.NewNodePool(10)

	for i := 0; i <= 10; i++ {
		addr := fmt.Sprintf("127.0.0.%d:26657", i+1)
		np.AddNode(addr)
		require.True(t, np.HasNode(addr))
	}
}

func TestNodePool_DeleteNode(t *testing.T) {
	np := crawl.NewNodePool(10)

	for i := 0; i <= 10; i++ {
		addr := fmt.Sprintf("127.0.0.%d:26657", i+1)
		np.AddNode(addr)
		require.True(t, np.HasNode(addr))
		np.DeleteNode(addr)
		require.False(t, np.HasNode(addr))
	}
}

func TestNodePool_Reseed(t *testing.T) {
	reseedSize := uint(10)
	np := crawl.NewNodePool(reseedSize)
	require.Equal(t, 0, np.Size())

	seeds := make([]string, 20)
	for i := range seeds {
		seeds[i] = fmt.Sprintf("127.0.0.%d:26657", i+1)
	}

	np.Seed(seeds)

	for _, s := range seeds {
		np.DeleteNode(s)
	}

	np.Reseed()
	require.Equal(t, reseedSize, uint(np.Size()))
}
