package server

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	gormlogger "gorm.io/gorm/logger"
)

var _ gormlogger.Interface = (*DBLogger)(nil)

// DBLogger implements a wrapper around the zerolog Logger type and implements
// the GORM logging interface.
type DBLogger struct {
	logger zerolog.Logger
}

// NewDBLogger creates a new wrapped DBLogger with a given zerolog logger.
func NewDBLogger(logger zerolog.Logger) *DBLogger {
	return &DBLogger{
		logger: logger.With().Str("module", "db").Logger(),
	}
}

// LogMode returns a new wrapped zerolog logger with a provided GORM log level
// which is mapped to a zerolog level.
func (dbl *DBLogger) LogMode(lvl gormlogger.LogLevel) gormlogger.Interface {
	var zLvl zerolog.Level
	switch lvl {
	case gormlogger.Silent:
		zLvl = zerolog.NoLevel

	case gormlogger.Error:
		zLvl = zerolog.ErrorLevel

	case gormlogger.Warn:
		zLvl = zerolog.WarnLevel

	case gormlogger.Info:
		zLvl = zerolog.InfoLevel
	}

	newLogger := *dbl
	newLogger.logger = dbl.logger.Level(zLvl)

	return &newLogger
}

func (dbl DBLogger) Info(_ context.Context, msg string, data ...interface{}) {
	dataStrs := make([]string, len(data))
	for i, d := range data {
		dataStrs[i] = fmt.Sprintf("%v", d)
	}

	dbl.logger.Info().Strs("data", dataStrs).Msg(msg)
}

func (dbl DBLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	dataStrs := make([]string, len(data))
	for i, d := range data {
		dataStrs[i] = fmt.Sprintf("%v", d)
	}

	dbl.logger.Warn().Strs("data", dataStrs).Msg(msg)
}

func (dbl DBLogger) Error(_ context.Context, msg string, data ...interface{}) {
	dataStrs := make([]string, len(data))
	for i, d := range data {
		dataStrs[i] = fmt.Sprintf("%v", d)
	}

	dbl.logger.Error().Strs("data", dataStrs).Msg(msg)
}

func (dbl DBLogger) Trace(_ context.Context, _ time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	if err != nil {
		dbl.logger.Error().Err(err).Int64("rows", rows).Str("sql", sql).Msg("")
	} else {
		dbl.logger.Info().Int64("rows", rows).Str("sql", sql).Msg("")
	}
}
