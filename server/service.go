package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/gologin/v2"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	httpswagger "github.com/swaggo/http-swagger"
	"golang.org/x/oauth2"
	githuboauth2 "golang.org/x/oauth2/github"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/cosmos/atlas/config"
	"github.com/cosmos/atlas/docs/api"
	v1 "github.com/cosmos/atlas/server/router/v1"
)

// @securityDefinitions.apikey APIKeyAuth
// @in header
// @name Authorization

// Service implements the encapsulating Atlas service. It wraps a router which is
// responsible for handling all versioned API requests with a given Router that
// interacts the with Atlas data model. The Service is responsible for establishing
// a database connection and managing session cookies.
type Service struct {
	logger       zerolog.Logger
	cfg          config.Config
	db           *gorm.DB
	cookieCfg    gologin.CookieConfig
	sessionStore *sessions.CookieStore
	oauth2Cfg    *oauth2.Config
	router       *mux.Router
	server       *http.Server
}

func NewService(logger zerolog.Logger, cfg config.Config) (*Service, error) {
	dbLogger := NewDBLogger(logger).LogMode(gormlogger.Silent)

	sessionKey, err := base64.StdEncoding.DecodeString(cfg.String(config.SessionKey))
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode session key: %w", err)
	}

	cookieCfg := gologin.DefaultCookieConfig
	sessionStore := sessions.NewCookieStore(sessionKey, nil)
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.Secure = true
	sessionStore.Options.MaxAge = 3600 * 24 * 7 // 1 week

	if cfg.Bool(config.Dev) {
		dbLogger = dbLogger.LogMode(gormlogger.Info)
		cookieCfg = gologin.DebugOnlyCookieConfig
		sessionStore.Options.Secure = false
	}

	db, err := gorm.Open(postgres.Open(cfg.String(config.DatabaseURL)), &gorm.Config{Logger: dbLogger})
	if err != nil {
		return nil, err
	}

	service := &Service{
		logger:       logger.With().Str("module", "server").Logger(),
		cfg:          cfg,
		db:           db,
		cookieCfg:    cookieCfg,
		sessionStore: sessionStore,
		router:       mux.NewRouter(),
		oauth2Cfg: &oauth2.Config{
			ClientID:     cfg.String(config.GHClientID),
			ClientSecret: cfg.String(config.GHClientSecret),
			Endpoint:     githuboauth2.Endpoint,
		},
	}

	v1Router, err := v1.NewRouter(service.logger, cfg, service.db, service.cookieCfg, service.sessionStore, service.oauth2Cfg)
	if err != nil {
		return nil, err
	}

	// mount api docs
	service.registerSwagger(cfg)

	// register v1 API routes
	v1Router.Register(service.router, v1.V1APIPathPrefix)

	return service, nil
}

// Start starts the atlas service as a blocking process.
func (s *Service) Start() error {
	s.server = &http.Server{
		Handler:      s.router,
		Addr:         s.cfg.String(config.ListenAddr),
		WriteTimeout: s.cfg.Duration(config.HTTPReadTimeout),
		ReadTimeout:  s.cfg.Duration(config.HTTPWriteTimeout),
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
		_ = s.server.Shutdown(ctx)
	}
}

func (s *Service) registerSwagger(cfg config.Config) {
	api.SwaggerInfo.Title = "Atlas API"
	api.SwaggerInfo.Description = "Atlas Cosmos SDK module registry API documentation."
	api.SwaggerInfo.Version = "1.0"
	api.SwaggerInfo.BasePath = v1.V1APIPathPrefix

	if cfg.Bool(config.Dev) {
		api.SwaggerInfo.Host = s.cfg.String(config.ListenAddr)
		api.SwaggerInfo.Schemes = []string{"http"}
	} else {
		api.SwaggerInfo.Host = "api.atlas.cosmos.network"
		api.SwaggerInfo.Schemes = []string{"https"}
	}

	// mount swagger API documentation
	s.router.PathPrefix("/api/docs").Handler(httpswagger.WrapHandler)
}
