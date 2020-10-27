package httputil_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/atlas/server/httputil"
)

func TestParsePaginationQueryParams(t *testing.T) {
	testCases := []struct {
		name      string
		req       *http.Request
		result    httputil.PaginationQuery
		expectErr bool
	}{
		{
			"valid next page",
			httptest.NewRequest("GET", "/foo?cursor=0&limit=10&page=next", nil),
			httputil.PaginationQuery{Cursor: "0", Limit: 10, Page: httputil.PageNext},
			false,
		},
		{
			"valid prev page",
			httptest.NewRequest("GET", "/foo?cursor=5&limit=10&page=prev", nil),
			httputil.PaginationQuery{Cursor: "5", Limit: 10, Page: httputil.PagePrev},
			false,
		},
		{
			"invalid cursor",
			httptest.NewRequest("GET", "/foo?cursor=-5&limit=10&page=prev", nil),
			httputil.PaginationQuery{},
			true,
		},
		{
			"invalid limit",
			httptest.NewRequest("GET", "/foo?cursor=0&limit=-10&page=next", nil),
			httputil.PaginationQuery{},
			true,
		},
		{
			"invalid page",
			httptest.NewRequest("GET", "/foo?cursor=0&limit=10&page=bat", nil),
			httputil.PaginationQuery{},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			pq, err := httputil.ParsePaginationQueryParams(tc.req)
			require.Equal(t, tc.expectErr, err != nil)
			require.Equal(t, tc.result, pq)
		})
	}
}
