package server

import (
	"net/http"

	"github.com/dghubble/sessions"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
	githuboauth2 "golang.org/x/oauth2/github"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/cosmos/atlas/config"
)

const (
	sessionName     = "example-github-app"
	sessionSecret   = "example cookie signing secret"
	sessionUserKey  = "githubID"
	sessionUsername = "githubUsername"
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
}

func NewService(logger zerolog.Logger, cfg config.Config) (*Service, error) {
	// TODO: Do we need to modify any GORM settings?
	db, err := gorm.Open(postgres.Open(cfg.String(config.FlagDatabaseURL)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	service := &Service{
		logger:     logger.With().Str("module", "server").Logger(),
		cfg:        cfg,
		router:     mux.NewRouter(),
		controller: NewController(db),
		oauth2Cfg: &oauth2.Config{
			ClientID:     cfg.String(config.FlagGHClientID),
			ClientSecret: cfg.String(config.FlagGHClientSecret),
			RedirectURL:  cfg.String(config.FlagGHRedirectURL),
			Endpoint:     githuboauth2.Endpoint,
		},
	}

	service.registerRoutes()
	return service, nil
}

func (s *Service) Start() error {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         s.cfg.String(config.FlagListenAddr),
		WriteTimeout: s.cfg.Duration(config.FlagHTTPReadTimeout),
		ReadTimeout:  s.cfg.Duration(config.FlagHTTPWriteTimeout),
	}

	s.logger.Info().Str("address", srv.Addr).Msg("starting atlas server...")
	return srv.ListenAndServe()
}

func (s *Service) registerRoutes() {
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

// func respondWithError(w http.ResponseWriter, code int, err error) {
// 	respondWithJSON(w, code, map[string]string{"error": err.Error()})
// }

// func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }
