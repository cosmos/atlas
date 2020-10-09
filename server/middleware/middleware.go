package middleware

import (
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

// Build returns a new middleware chain.
func Build(logger zerolog.Logger) alice.Chain {
	mChain := alice.New()
	mChain = AddRequestLoggingMiddleware(mChain, logger)

	return mChain
}

// AddRequestLoggingMiddleware returns an HTTP request logging middleware.
func AddRequestLoggingMiddleware(mChain alice.Chain, logger zerolog.Logger) alice.Chain {
	mChain = mChain.Append(hlog.NewHandler(logger))
	mChain = mChain.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))
	mChain = mChain.Append(hlog.RequestHandler("req"))
	mChain = mChain.Append(hlog.RemoteAddrHandler("ip"))
	mChain = mChain.Append(hlog.UserAgentHandler("ua"))
	mChain = mChain.Append(hlog.RefererHandler("ref"))
	mChain = mChain.Append(hlog.RequestIDHandler("req_id", "Request-Id"))

	return mChain
}
