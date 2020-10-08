package config

import (
	"time"
)

// Configuration values that may be provided in a configuration file, as
// environment variables or via CLI flags. Note, not all configurations may be
// passed as CLI flags. All keys are dot delimitated except for environment
// variables which are snake-cased and must be prefixed with ATLAS_*.
const (
	FlagConfig           = "config"
	FlagLogLevel         = "log.level"
	FlagLogFormat        = "log.format"
	FlagListenAddr       = "listen.addr"
	FlagDev              = "dev"
	FlagDatabaseURL      = "database.url"
	FlagHTTPReadTimeout  = "http.read.timeout"
	FlagHTTPWriteTimeout = "http.write.timeout"
	FlagGHClientID       = "gh.client.id"
	FlagGHClientSecret   = "gh.client.secret"
	FlagGHRedirectURL    = "gh.redirect.url"
	FlagSessionKey       = "session.key"
)

// Config defines a configuration abstraction so we don't rely on any specific
// concrete configuration manager.
type Config interface {
	Bool(path string) bool
	String(path string) string
	Int(path string) int
	Ints(path string) []int
	Duration(path string) time.Duration
}
