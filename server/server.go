package server

import (
	"context"
	"net/http"
	"time"

	"github.com/dghubble/sessions"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"golang.org/x/oauth2"
	githuboauth2 "golang.org/x/oauth2/github"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/cosmos/atlas/config"
)

const (
	sessionName     = "example-github-app"
	sessionSecret   = "example cookie signing secret"
	sessionUserKey  = "githubID"
	sessionUsername = "githubUsername"

	methodGET = "GET"
)

var sessionStore = sessions.NewCookieStore([]byte(sessionSecret), nil)

// Service defines the encapsulating Atlas service. It wraps a router which is
// responsible for handling API requests with a given controller that interacts
// with Atlas models. The Service is responsible for establishing a database
// connection and managing session cookies.
type Service struct {
	logger     zerolog.Logger
	cfg        config.Config
	router     *mux.Router
	controller *Controller
	oauth2Cfg  *oauth2.Config
	server     *http.Server
	db         *gorm.DB
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

	db.Logger = dbLogger
	service := &Service{
		logger:     logger.With().Str("module", "server").Logger(),
		cfg:        cfg,
		router:     mux.NewRouter(),
		db:         db,
		controller: NewController(db),
		oauth2Cfg: &oauth2.Config{
			ClientID:     cfg.String(config.FlagGHClientID),
			ClientSecret: cfg.String(config.FlagGHClientSecret),
			RedirectURL:  cfg.String(config.FlagGHRedirectURL),
			Endpoint:     githuboauth2.Endpoint,
		},
	}

	service.registerV1Routes()
	return service, nil
}

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

	// TODO: Consider enabling rate limiting.

	return mChain
}

func (s *Service) registerV1Routes() {
	// Create a versioned sub-router. All routes will be mounted under this
	// sub-router.
	v1Router := s.router.PathPrefix("/api/v1").Subrouter()

	// build middleware chain
	mChain := s.buildMiddleware()

	// unauthenticated routes
	v1Router.Handle("/modules/{id}", mChain.ThenFunc(s.controller.GetModuleByID())).Methods(methodGET)

	// authenticated routes

	// session routes

	// // middleware
	// v1Router.Use(NewLogRequestMiddleware(s.logger).Handle)

	// s.router.HandleFunc("/", homeHandler)
	// s.router.HandleFunc("/logout", logoutHandler)

	// // state param cookies require HTTPS by default; disable for localhost development
	// stateConfig := gologin.DebugOnlyCookieConfig
	// s.router.Handle("/github/login", github.StateHandler(stateConfig, github.LoginHandler(oauth2Config, nil)))
	// s.router.Handle("/github/callback", github.StateHandler(stateConfig, github.CallbackHandler(oauth2Config, issueSession(), nil)))
}

// func issueSession() http.Handler {
// 	fn := func(w http.ResponseWriter, req *http.Request) {
// 		ctx := req.Context()
// 		githubUser, err := github.UserFromContext(ctx)
// 		if err != nil {
// 			respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get session: %w", err))
// 			return
// 		}

// 		// 2. Implement a success handler to issue some form of session
// 		session := sessionStore.New(sessionName)
// 		session.Values[sessionUserKey] = *githubUser.ID
// 		session.Values[sessionUsername] = *githubUser.Login
// 		session.Save(w)

// 		http.Redirect(w, req, "/", http.StatusFound)
// 	}

// 	return http.HandlerFunc(fn)
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	session, err := sessionStore.Get(r, sessionName)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, fmt.Errorf("failed to get session: %w", err))
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, session)
// }

// func logoutHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "POST" {
// 		fmt.Println("DESTRYOING ESSION")
// 		sessionStore.Destroy(w, sessionName)
// 	}

// 	http.Redirect(w, req, "/", http.StatusFound)
// }
