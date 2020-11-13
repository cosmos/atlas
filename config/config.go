package config

import (
	"time"
)

// Configuration values that may be provided in a configuration file, as
// environment variables or via CLI flags. Note, not all configurations may be
// passed as CLI flags. All keys are dot delimitated except for environment
// variables which are snake-cased and must be prefixed with ATLAS_*.
const (
	ConfigPath       = "config"
	LogLevel         = "log.level"
	LogFormat        = "log.format"
	ListenAddr       = "listen.addr"
	Dev              = "dev"
	DatabaseURL      = "database.url"
	HTTPReadTimeout  = "http.read.timeout"
	HTTPWriteTimeout = "http.write.timeout"
	GHClientID       = "gh.client.id"
	GHClientSecret   = "gh.client.secret"
	SessionKey       = "session.key"
	AllowedOrigins   = "allowed.origins"
	SendGridAPIKey   = "sendgrid.api.key"
	DomainName       = "domain.name"
	SyslogAddr       = "syslog.addr"
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
