package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/cosmos/atlas/config"
)

const (
	methodGET  = "GET"
	methodPOST = "POST"
)

// Service defines the encapsulating Atlas service. It wraps a router which is
// responsible for handling API requests with a given controller that interacts
// with Atlas models. The Service is responsible for establishing a database
// connection and managing session cookies.
type Service struct {
	logger     zerolog.Logger
	cfg        config.Config
	db         *gorm.DB
	router     *mux.Router
	controller *Controller
	server     *http.Server
}

func NewService(logger zerolog.Logger, cfg config.Config) (*Service, error) {
	dbLogger := NewDBLogger(logger).LogMode(gormlogger.Silent)
	if cfg.Bool(config.FlagDev) {
		dbLogger = dbLogger.LogMode(gormlogger.Info)
	}

	db, err := gorm.Open(postgres.Open(cfg.String(config.FlagDatabaseURL)), &gorm.Config{Logger: dbLogger})
	if err != nil {
		return nil, err
	}

	service := &Service{
		logger:     logger.With().Str("module", "server").Logger(),
		cfg:        cfg,
		db:         db,
		router:     mux.NewRouter(),
		controller: NewController(db, cfg),
	}

	service.registerV1Routes()
	return service, nil
}

// Start starts the atlas service as a blocking process.
func (s *Service) Start() error {
	s.server = &http.Server{
		Handler:      s.router,
		Addr:         s.cfg.String(config.FlagListenAddr),
		WriteTimeout: s.cfg.Duration(config.FlagHTTPReadTimeout),
		ReadTimeout:  s.cfg.Duration(config.FlagHTTPWriteTimeout),
	}

	s.logger.Info().Str("address", s.server.Addr).Msg("starting atlas server...")
	return s.server.ListenAndServe()
}

// Cleanup performs server cleanup. If the internal HTTP server is non-nil, the
// server will be shutdown after a grace period deadline.
func (s *Service) Cleanup() {
	if s.server != nil {
		// create a deadline to wait for all existing requests to finish
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Do not block if no connections exist, but otherwise, we will wait until
		// the timeout deadline.
		s.server.Shutdown(ctx)
	}
}

func (s *Service) buildMiddleware() alice.Chain {
	mChain := alice.New()
	mChain = mChain.Append(hlog.NewHandler(s.logger))
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

func (s *Service) registerV1Routes() {
	// Create a versioned sub-router. All API routes will be mounted under this
	// sub-router.
	v1Router := s.router.PathPrefix("/api/v1").Subrouter()

	// build middleware chain
	mChain := s.buildMiddleware()

	// unauthenticated routes
	v1Router.Handle(
		"/modules",
		mChain.ThenFunc(s.controller.GetAllModules()),
	).Queries("cursor", "{cursor:[0-9]+}", "limit", "{limit:[0-9]+}").Methods(methodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}",
		mChain.ThenFunc(s.controller.GetModuleByID()),
	).Methods(methodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}/versions",
		mChain.ThenFunc(s.controller.GetModuleVersions()),
	).Methods(methodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}/authors",
		mChain.ThenFunc(s.controller.GetModuleAuthors()),
	).Methods(methodGET)

	v1Router.Handle(
		"/modules/{id:[0-9]+}/keywords",
		mChain.ThenFunc(s.controller.GetModuleKeywords()),
	).Methods(methodGET)

	v1Router.Handle(
		"/users/{id:[0-9]+}",
		mChain.ThenFunc(s.controller.GetUserByID()),
	).Methods(methodGET)

	v1Router.Handle(
		"/users/{id:[0-9]+}/modules",
		mChain.ThenFunc(s.controller.GetUserModules()),
	).Methods(methodGET)

	v1Router.Handle(
		"/users",
		mChain.ThenFunc(s.controller.GetAllUsers()),
	).Queries("cursor", "{cursor:[0-9]+}", "limit", "{limit:[0-9]+}").Methods(methodGET)

	v1Router.Handle(
		"/keywords",
		mChain.ThenFunc(s.controller.GetAllKeywords()),
	).Queries("cursor", "{cursor:[0-9]+}", "limit", "{limit:[0-9]+}").Methods(methodGET)

	// authenticated routes

	// session routes
	v1Router.Handle(
		"/session/start",
		mChain.Then(s.controller.BeginSession()),
	).Methods(methodGET)

	v1Router.Handle(
		"/session/authorize",
		mChain.Then(s.controller.AuthorizeSession()),
	).Methods(methodGET)

	v1Router.Handle(
		"/session/logout",
		mChain.ThenFunc(s.controller.LogoutSession()),
	).Methods(methodPOST)
}
