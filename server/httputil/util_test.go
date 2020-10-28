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
			httptest.NewRequest("GET", "/foo?page=1&limit=10", nil),
			httputil.PaginationQuery{Page: 1, Limit: 10, Order: "id", Reverse: false},
			false,
		},
		{
			"valid next page with order",
			httptest.NewRequest("GET", "/foo?page=1&limit=10&order=created_at", nil),
			httputil.PaginationQuery{Page: 1, Limit: 10, Order: "created_at,id", Reverse: false},
			false,
		},
		{
			"valid next page with order and reverse",
			httptest.NewRequest("GET", "/foo?page=1&limit=10&order=created_at&reverse=true", nil),
			httputil.PaginationQuery{Page: 1, Limit: 10, Order: "created_at,id", Reverse: true},
			false,
		},
		{
			"invalid non-positive page",
			httptest.NewRequest("GET", "/foo?page=0&limit=10", nil),
			httputil.PaginationQuery{},
			true,
		},
		{
			"invalid non-numeric page",
			httptest.NewRequest("GET", "/foo?page=bar&limit=10", nil),
			httputil.PaginationQuery{},
			true,
		},
		{
			"invalid non-positive limit",
			httptest.NewRequest("GET", "/foo?page=1&limit=-10", nil),
			httputil.PaginationQuery{},
			true,
		},
		{
			"invalid non-numeric limit",
			httptest.NewRequest("GET", "/foo?page=1&limit=bar", nil),
			httputil.PaginationQuery{},
			true,
		},
		{
			"invalid reverse",
			httptest.NewRequest("GET", "/foo?page=1&limit=10&order=created_at&reverse=bar", nil),
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
