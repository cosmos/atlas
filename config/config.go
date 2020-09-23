package config

import "time"

// Configuration values that may be provided in a configuration file, as
// environment variables or via CLI flags.
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
)

// Config defines a configuration abstraction provided to a Server type in order
// to be instantiated.
type Config interface {
	String(path string) string
	Int(path string) int
	Ints(path string) []int
	Duration(path string) time.Duration
}
