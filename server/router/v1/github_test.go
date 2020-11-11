package v1

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGitHubClient_GetRepository(t *testing.T) {
	// Note: An API access key is not necessarily required to execute this test,
	// however, API limits are easily reached if executed multiple times in a short
	// timeframe.
	client := NewGitHubClient(os.Getenv("ATLAS_TEST_GITHUB_ACCESS_KEY"))

	testCases := []struct {
		name        string
		repoURL     string
		expectOwner string
		expectRepo  string
		expectErr   bool
	}{
		{
			"invalid URL",
			"foo",
			"", "",
			true,
		},
		{
			"invalid repo",
			"https://github.com/cosmos",
			"", "",
			true,
		},
		{
			"valid repo with extension",
			"https://github.com/cosmos/cosmos-sdk.git",
			"cosmos", "cosmos-sdk",
			false,
		},
		{
			"valid repo with query",
			"https://github.com/cosmos/cosmos-sdk?foo=bar",
			"cosmos", "cosmos-sdk",
			false,
		},
		{
			"valid repo with extension and query",
			"https://github.com/cosmos/cosmos-sdk.git?foo=bar",
			"cosmos", "cosmos-sdk",
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			repo, err := client.GetRepository(tc.repoURL)
			if tc.expectErr {
				require.Error(t, err)
				require.Empty(t, repo)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectOwner, repo.Owner)
				require.Equal(t, tc.expectRepo, repo.Repo)
				require.Greater(t, len(repo.Contributors), 0)
			}
		})
	}
}
