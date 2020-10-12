package middleware

import (
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"

	"github.com/cosmos/atlas/config"
	"github.com/cosmos/atlas/server/httputil"
)

// Build returns a new middleware chain.
func Build(logger zerolog.Logger, cfg config.Config) alice.Chain {
	mChain := alice.New()
	mChain = AddRequestLoggingMiddleware(mChain, logger)
	mChain = AddCORSMiddleware(mChain, logger, cfg)

	return mChain
}

// AddRequestLoggingMiddleware appends HTTP logging middleware to a provided
// middleware chain.
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

// AddCORSMiddleware appends CORS middleware to a provided middleware chain.
func AddCORSMiddleware(mChain alice.Chain, logger zerolog.Logger, cfg config.Config) alice.Chain {
	c := cors.New(cors.Options{
		AllowedMethods: []string{
			httputil.MethodGET,
			httputil.MethodGET,
			httputil.MethodPOST,
			httputil.MethodPUT,
			httputil.MethodDELETE,
		},
		AllowCredentials: true,
		AllowedOrigins:   []string{"https://atlas.cosmos.network"},
		AllowedHeaders:   []string{"*"},
		Debug:            false,
	})

	if cfg.Bool(config.FlagDev) {
		c = cors.AllowAll()
	}

	c.Log = &logger
	mChain = mChain.Append(c.Handler)

	return mChain
}
