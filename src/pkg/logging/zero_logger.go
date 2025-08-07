package logging

import (
	"os"

	"github.com/amirhasanpour/bama-clone-web-api/src/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var zeroSinLogger *zerolog.Logger

type zeroLogger struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

var zeroLogLevelMapping = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func newZeroLogger(cfg *config.Config) *zeroLogger {
	logger := &zeroLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zeroLogger) getLogLevel() zerolog.Level {
	level, exists := zeroLogLevelMapping[l.cfg.Logger.Level]
	if !exists {
		return zerolog.DebugLevel
	}
	return level
}

func (l *zeroLogger) Init() {
	once.Do(func() {

		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

		file, err := os.OpenFile(l.cfg.Logger.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic("could not open log file")
		}

		var logger = zerolog.New(file).
			With().
			Timestamp().
			Str("AppName", "MyApp").
			Str("LoggerName", "Zerolog").
			Logger()
		zerolog.SetGlobalLevel(l.getLogLevel())
		zeroSinLogger = &logger
	})
	l.logger = zeroSinLogger
}

func (l *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {

	l.logger.
		Debug().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Debugf(template string, args ...any) {
	l.logger.
		Debug().
		Msgf(template, args...)
}

func (l *zeroLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {

	l.logger.
		Info().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Infof(template string, args ...any) {
	l.logger.
		Info().
		Msgf(template, args...)
}

func (l *zeroLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {

	l.logger.
		Warn().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Warnf(template string, args ...any) {
	l.logger.
		Warn().
		Msgf(template, args...)
}

func (l *zeroLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {

	l.logger.
		Error().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Errorf(template string, args ...any) {
	l.logger.
		Error().
		Msgf(template, args...)
}

func (l *zeroLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {

	l.logger.
		Fatal().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Fatalf(template string, args ...any) {
	l.logger.
		Fatal().
		Msgf(template, args...)
}
